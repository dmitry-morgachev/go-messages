package service

import (
	"log"
	"message-service/internal/db"
)

func SaveMessage(content string) error {
	query := `INSERT INTO messages (content) VALUES ($1)`
	_, err := db.DB.Exec(query, content)
	if err != nil {
		log.Printf("Error inserting message: %v", err)
		return err
	}

	return nil
}

func MarkMessageAsProcessed(content string) error {
	query := `UPDATE messages SET is_processed = true, processed_at = NOW() WHERE content = $1`
	_, err := db.DB.Exec(query, content)
	if err != nil {
		log.Printf("Error marking message as processed: %v", err)
		return err
	}
	return nil
}

func GetProcessedStats() (int, []string, error) {
	var totalProcessed int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM messages WHERE is_processed = true`).Scan(&totalProcessed)
	if err != nil {
		return 0, nil, err
	}

	rows, err := db.DB.Query(`SELECT content FROM messages WHERE is_processed = true ORDER BY processed_at DESC LIMIT 10`)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	var lastProcessed []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return 0, nil, err
		}
		lastProcessed = append(lastProcessed, content)
	}

	return totalProcessed, lastProcessed, nil
}
