package main

import (
	"net/http"
	"time"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:     false,
		Message:   "hit the broker-service",
		CreatedAt: time.Now(),
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Root(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:     false,
		Message:   "broker-service is running",
		CreatedAt: time.Now(),
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
