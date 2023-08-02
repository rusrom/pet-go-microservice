package main

import (
	"net/http"
	"time"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:     false,
		Message:   "Hit the broker",
		CreatedAt: time.Now(),
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
