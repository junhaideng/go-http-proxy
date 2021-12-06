package main

import (
	"log"

	"github.com/junhaideng/go-proxy/config"
	"github.com/junhaideng/go-proxy/proxy"
)

func main(){
	cfg, err := config.New("config/config.yml")
	if err != nil {
		panic(err)
	}

	s := proxy.NewServer(cfg)
	log.Printf("Start server on: %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	s.Run()
}