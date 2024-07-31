package handler

import (
	"encoding/json"
	"message-service/internal/kafka"
	"message-service/internal/model"
	"message-service/internal/service"
	"net/http"
)

func SaveMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message model.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := service.SaveMessage(message.Content); err != nil {
		http.Error(w, "Could not save message", http.StatusInternalServerError)
		return
	}

	if err := kafka.ProduceMessage("messages", message.Content); err != nil {
		http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
