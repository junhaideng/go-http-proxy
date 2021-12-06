package proxy

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/junhaideng/go-proxy/config"
	"github.com/junhaideng/go-proxy/http"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		config: cfg,
	}

	return s
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept connection failed: ", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	request, err := http.ParseRequest(conn)
	if err != nil {
		conn.Close()
		log.Println("Parse requst failed")
		return
	}

	// 检查校验
	if s.config.Server.EnableAuth {
		auth := s.config.Server.Auth
		if !request.ProxyAuth().Check(auth.Username, auth.Password) {
			http.RequireProxyAuth(conn)
			return
		}
	}

	host, ok := request.Host()
	if !ok {
		conn.Close()
		log.Println("No Host specified")
		return
	}

	if !strings.Contains(host, ":") {
		host += ":80"
	}

	server, err := net.Dial("tcp", host)
	if err != nil {
		conn.Close()
		log.Println("Dial server failed: ", err)
		return
	}

	if request.Method == "CONNECT" {
		log.Println("Visit: ", request.Path)
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		tunnel(conn, server)
		return
	}

	log.Println("Visit: ", request.Path)

	_, err = server.Write(request.Raw())
	if err != nil {
		log.Println("Write server failed: ", err)
		conn.Close()
		server.Close()
		return
	}

	tunnel(conn, server)
}
