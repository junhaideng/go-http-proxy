package http

import (
	"strings"
	"testing"
)

func TestParseRequestWithoutBody(t *testing.T) {
	r := strings.NewReader("GET / HTTP/1.1\r\nHost:httpbin.org\r\n\r\n")
	request, err := ParseRequest(r)
	if err != nil {
		t.Fatal(err)
	}
	if request.Method != "GET" {
		t.Fatal("Method is not get")
	}
	if request.Path != "/" {
		t.Fatal("Path is not /")
	}
	if request.Version != "HTTP/1.1" {
		t.Fatal("HTTP version is not right")
	}

	host := request.Headers.Get("Host")
	if host == "" {
		t.Fatal("Host is empty")
	}
}

func TestParseRequestWithBody(t *testing.T) {
	r := strings.NewReader("POST / HTTP/1.1\r\nHost:httpbin.org\r\nContent-Length:9\r\n\r\nuser=test")
	request, err := ParseRequest(r)
	if err != nil {
		t.Fatal(err)
	}
	if request.Method != "POST" {
		t.Fatal("Method is not get")
	}
	if request.Path != "/" {
		t.Fatal("Path is not /")
	}
	if request.Version != "HTTP/1.1" {
		t.Fatal("HTTP version is not right")
	}

	host := request.Headers.Get("Host")
	if host == "" {
		t.Fatal("Host is empty")
	}

	body := request.Body
	if len(body) != 9 {
		t.Fatal("Body Length is not right")
	}
	if string(body) != "user=test" {
		t.Fatal("Body is not matched")
	}
}
