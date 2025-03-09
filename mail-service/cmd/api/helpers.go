package main

//ovo se moze odvojiti u poseban pkg, pa da se importuje, imamo vec 2 ista ovakva fajla
import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	doc := json.NewDecoder(r.Body)
	err := doc.Decode(data)
	// s obzirom da je data interface, nema mu potrebe izgleda prosledjivati ampersand
	if err != nil {
		return err
	}

	err = doc.Decode(&struct{}{})
	// ova dodjela sluzi da se provjeri da li ima vise od
	// jednog json objekta proslijedjenog, zato sada provjeravamo
	// da li je error zapravo io.EOF (end of file)
	if err != io.EOF {
		return errors.New("body must contain only one JSON object ")
	}

	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, val := range headers[0] {
			w.Header()[key] = val
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

func (app *Config) errJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse

	payload.Error = true
	payload.Message = err.Error()
	return app.writeJSON(w, statusCode, payload)
}
