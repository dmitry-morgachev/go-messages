package handler

import (
	"encoding/json"
	"message-service/internal/service"
	"net/http"
)

type StatsResponse struct {
	TotalProcessed int      `json:"total_processed"`
	LastProcessed  []string `json:"last_processed"`
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получение статистики из сервиса
	totalProcessed, lastProcessed, err := service.GetProcessedStats()
	if err != nil {
		http.Error(w, "Failed to get statistics", http.StatusInternalServerError)
		return
	}

	response := StatsResponse{
		TotalProcessed: totalProcessed,
		LastProcessed:  lastProcessed,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
