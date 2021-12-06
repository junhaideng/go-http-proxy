package http

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/junhaideng/go-proxy/internal/byteconv"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Headers http.Header
	Body    []byte
	raw     []byte
}

func (r Request) Host() (string, bool) {
	if r.Headers.Get("Host") != "" {
		return r.Headers.Get("Host"), true
	}
	return "", false
}

func (r Request) Raw() []byte {
	return r.raw
}

func (r Request) ProxyAuth() *Auth {
	p := r.Headers.Get("Proxy-Authorization")
	if p == "" {
		return nil
	}
	auth, err := parseAuth(p)
	if err != nil {
		return nil
	}
	return auth
}

func (r Request) BasicAuth() *Auth {
	b := r.Headers.Get("Basic-Authorization")
	if b == "" {
		return zeroAuth
	}
	auth, err := parseAuth(b)
	if err != nil {
		return zeroAuth
	}

	return auth
}


func ParseRequest(conn io.Reader) (*Request, error) {
	br := bufio.NewReader(conn)

	// 第一行
	line, err := br.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	if bytes.Count(line, []byte{' '}) < 2 {
		return nil, errors.New("")
	}
	// 第一部分为请求方法
	i := bytes.IndexByte(line, ' ')

	j := bytes.LastIndexByte(line, ' ')

	req := &Request{Headers: make(http.Header), raw: make([]byte, 0, 1024)}
	req.raw = append(req.raw, line...)

	req.Method = byteconv.Bytes2Str(line[:i])
	req.Path = byteconv.Bytes2Str(line[i+1 : j])
	req.Version = byteconv.Bytes2Str(bytes.TrimSpace(line[j+1:]))

	// 解析请求头部
	for {
		line, err := br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		req.raw = append(req.raw, line...)
		line = bytes.TrimSpace(line)
		// \r\n
		if len(line) == 0 {
			break
		}
		colon := bytes.IndexByte(line, ':')
		req.Headers.Add(byteconv.Bytes2Str(bytes.TrimSpace(line[:colon])), byteconv.Bytes2Str(bytes.TrimSpace(line[colon+1:])))
	}

	contentLength := req.Headers.Get("Content-Length")
	if contentLength == "" {
		return req, nil
	}

	length, err := strconv.Atoi(contentLength)
	if err != nil {
		return nil, err
	}

	body := make([]byte, length)
	n, err := br.Read(body)
	if err != nil {
		return nil, err
	}
	if n != length {
		return nil, errors.New("body is not equal to Content-Length")
	}
	req.raw = append(req.raw, body...)

	// 读取 body
	req.Body = body

	return req, nil
}


