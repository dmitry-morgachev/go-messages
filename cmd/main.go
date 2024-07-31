package main

import (
	"log"
	"message-service/internal/db"
	"message-service/internal/handler"
	"message-service/internal/kafka"
	"message-service/pkg/config"
	"net/http"
)

func main() {
	cfg := config.GetConfig()

	db.InitDB(cfg)
	kafka.InitKafka(cfg)

	go kafka.StartConsumer(cfg)

	http.HandleFunc("/message/save", handler.SaveMessageHandler)
	http.HandleFunc("/stats", handler.StatsHandler)

	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
