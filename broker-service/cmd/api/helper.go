package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

// readJson reads the request body and
func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	const maxBytes int64 = 1024 * 1024 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	dummy := &struct{}{}
	err = dec.Decode(dummy)
	if err != io.EOF {
		return errors.New("only a single value is accepted in the JSON body")
	}

	return nil
}

func (app *Config) writeJson(
	w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) writeError(w http.ResponseWriter, err error, status ...int) error {
	payload := jsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	return app.writeJson(w, statusCode, payload)
}
