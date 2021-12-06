package main

import (
	"flag"
	"log"
	"os"

	"github.com/junhaideng/go-proxy/config"
	"github.com/junhaideng/go-proxy/proxy"
)

var conf string
var logPath string

func init() {
	flag.StringVar(&conf, "c", "config/config.yml", "config file path")
	flag.StringVar(&logPath, "w", "", "write log to file")
}

func main() {
	flag.Parse()
	cfg, err := config.New(conf)
	if err != nil {
		panic(err)
	}

	if logPath != "" {
		f, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		log.SetOutput(f)
	}

	s := proxy.NewServer(cfg)
	log.Printf("Start server on: %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	s.Run()
}
