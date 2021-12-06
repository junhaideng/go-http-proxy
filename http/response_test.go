package http

import (
	"strings"
	"testing"
)

func TestParseResponseWithoutBody(t *testing.T) {
	r := strings.NewReader("HTTP/1.1 200 OK\r\nHost:httpbin.org\r\n\r\n")
	response, err := ParseResponse(r)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 200 {
		t.Fatal("Method is not get")
	}
	if response.StatusText != "OK" {
		t.Fatal("Path is not /")
	}
	if response.Version != "HTTP/1.1" {
		t.Fatal("HTTP version is not right")
	}

	host := response.Headers.Get("Host")
	if host == "" {
		t.Fatal("Host is empty")
	}
}

func TestParseResponseWithBody(t *testing.T) {
	r := strings.NewReader("HTTP/1.1 200 OK\r\nHost:httpbin.org\r\nContent-Length:9\r\n\r\nuser=test")
	response, err := ParseResponse(r)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 200 {
		t.Fatal("Method is not get")
	}
	if response.StatusText != "OK" {
		t.Fatal("Path is not /")
	}
	if response.Version != "HTTP/1.1" {
		t.Fatal("HTTP version is not right")
	}

	host := response.Headers.Get("Host")
	if host == "" {
		t.Fatal("Host is empty")
	}

	body := response.Body
	if len(body) != 9 {
		t.Fatal("Body Length is not right")
	}
	if string(body) != "user=test" {
		t.Fatal("Body is not matched")
	}
}
