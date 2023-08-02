package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type jsonResponse struct {
	Error     bool      `json:"error"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	const maxBytes = 1048576 // 1 Mb

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	if err = dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
func (app *Config) writeJSON(w http.ResponseWriter, status int, data any) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	//if len(headers) > 0 {
	//	for key, value := range headers {
	//		w.Header()[key] = value
	//	}
	//}

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, statusCode int) error {
	data := jsonResponse{
		Error:   false,
		Message: err.Error(),
	}
	return app.writeJSON(w, statusCode, data)
}
