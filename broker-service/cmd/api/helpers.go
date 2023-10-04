package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// sample jsonresponse
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1 mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(int64(maxBytes)))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{}) // check if equal 2 bracket
	if err != io.EOF {
		return errors.New("body must have only a single JSON Value")
	}

	return nil
}
