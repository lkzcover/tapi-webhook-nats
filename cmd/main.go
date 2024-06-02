package main

import (
	"github.com/lkzcover/webhook-nats/internal/pkg/config"
	"github.com/lkzcover/webhook-nats/internal/pkg/http"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	nc, err := nats.Connect(cfg.NatsConn)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting HTTPS server on localhost:%s", cfg.HttpServer.Port)
	httpsServer := http.NewServer(nc, cfg.NastSubj, true)
	err = httpsServer.Run(cfg.HttpServer.Port, cfg.HttpServer.Cert, cfg.HttpServer.Key)
	if err != nil {
		log.Fatal(err)
	}
}
