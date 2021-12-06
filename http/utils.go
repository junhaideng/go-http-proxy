package http

import (
	"net"

	"github.com/junhaideng/go-proxy/internal/byteconv"
)

func Forbidden(conn net.Conn) {
	conn.Write(byteconv.Str2Byte(("HTTP/1.1 403 Forbidden\r\n\r\n")))
	conn.Close()
}

func OK(conn net.Conn) {
	conn.Write(byteconv.Str2Byte("HTTP/1.1 200 OK\r\n\r\n"))
	conn.Close()
}

func Unauthorized(conn net.Conn) {
	conn.Write(byteconv.Str2Byte("HTTP/1.1 401 Unauthorized\r\n\r\n"))
	conn.Close()
}

func NotFound(conn net.Conn) {
	conn.Write(byteconv.Str2Byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	conn.Close()
}

func RequireProxyAuth(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 407 Proxy Authentication Required\r\n\r\n"))
	conn.Close()
}
