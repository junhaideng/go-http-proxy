package http

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/junhaideng/go-proxy/internal/byteconv"
)

type Response struct {
	Version    string
	StatusCode int
	StatusText string
	Headers    http.Header
	Body       []byte
	raw        []byte
}

func ParseResponse(conn io.Reader) (*Response, error) {
	br := bufio.NewReader(conn)

	// 第一行
	line, err := br.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	if bytes.Count(line, []byte{' '}) < 2 {
		return nil, errors.New("")
	}

	line = bytes.TrimSpace(line)

	// 第一部分为请求方法
	i := bytes.IndexByte(line, ' ')

	j := bytes.LastIndexByte(line, ' ')

	resp := &Response{Headers: make(http.Header), raw: make([]byte, 0, 1024)}
	resp.raw = append(resp.raw, line...)

	resp.Version = byteconv.Bytes2Str(line[:i])
	code := byteconv.Bytes2Str(line[i+1 : j])
	resp.StatusCode, err = strconv.Atoi(code)
	if err != nil {
		return nil, fmt.Errorf("status code: %s is not an integer", code)
	}

	resp.StatusText = byteconv.Bytes2Str(bytes.TrimSpace(line[j+1:]))

	for {
		line, err := br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		resp.raw = append(resp.raw, line...)
		line = bytes.TrimSpace(line)
		// \r\n
		if len(line) == 0 {
			break
		}
		colon := bytes.IndexByte(line, ':')
		resp.Headers.Add(byteconv.Bytes2Str(bytes.TrimSpace(line[:colon])), byteconv.Bytes2Str(bytes.TrimSpace(line[colon+1:])))
	}

	contentLength := resp.Headers.Get("Content-Length")
	if contentLength == "" {
		return resp, nil
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

	resp.raw = append(resp.raw, body...)

	// 读取 body
	resp.Body = body

	return resp, nil
}
