package proxy

import (
	"io"
	"net"
)

// http tunnel
func tunnel(client net.Conn, server net.Conn) {
	go io.Copy(server, client)
	go io.Copy(client, server)
}

