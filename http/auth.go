package http

import (
	"bytes"
	"encoding/base64"
	"errors"

	"github.com/junhaideng/go-proxy/internal/byteconv"
)

var zeroAuth *Auth = &Auth{}

type Auth struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (a Auth) Check(username, password string) bool {
	return username == a.Username && password == a.Password
}

func parseAuth(s string) (*Auth, error) {
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	result := bytes.Split(res, []byte{':'})
	if len(result) != 2 {
		return nil, errors.New("string is invalid")
	}

	return &Auth{
		Username: byteconv.Bytes2Str(result[0]),
		Password: byteconv.Bytes2Str(result[1]),
	}, nil
}
