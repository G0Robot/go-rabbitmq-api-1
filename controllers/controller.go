package controllers

import (
	"encoding/json"
	"go-rabbitmq-api-1/models"
	"go-rabbitmq-api-1/services"
	"net/http"
)

// POST SendMessage
func SendMessage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var message models.Message
	_ = json.NewDecoder(r.Body).Decode(&message)

	services.SendMessageService(message)
	json.NewEncoder(w).Encode("Message Sent.")
}
