package http

import (
	"encoding/base64"
	"testing"
)

func TestParseAuth(t *testing.T) {
	s := base64.StdEncoding.EncodeToString([]byte("username:password"))

	a, err := parseAuth(s)
	if err != nil {
		t.Fatal(err)
	}

	if a.Username != "username" || a.Password != "password" {
		t.Fatal()
	}
}
