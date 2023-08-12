// Package payload provides utilities for dealing with HTTP request and response payloads.
// It integrates with sibling packages log and errors.
package utils

import (
	"compress/gzip"
	"ecommerce/errors"
	"encoding/json"
	"log"
	"net/http"
)

var (
	encodedErrResp []byte = json.RawMessage(`{"message":"Something has gone wrong"}`)
)

// ClientReporter provides information about an error such that client and
// server errors can be distinguished and handled appropriately.
type ClientReporter interface {
	error
	Message() map[string]string
	Status() int
}

// WriteError writes an appropriate error response to the given response
// writer. If the given error implements ClientReport, then the values from
// ErrorReport() and StatusCode() are written to the response, except in
// the case of a 5XX error, where the error is logged and a default message is
// written to the response.
func WriteError(w http.ResponseWriter, r *http.Request, e error) {
	if cr, ok := e.(ClientReporter); ok {
		status := cr.Status()
		if status >= http.StatusInternalServerError {
			handleInternalServerError(w, r, e)
			return
		}

		// log.FromRequest(r).Print(cr.Error())
		Write(w, r, cr.Message(), status)

		return
	}

	handleInternalServerError(w, r, e)
}

// Write writes the given payload to the response. If the payload
// cannot be marshaled, a 500 error is written instead. If the writer
// cannot be written to, then this function panics.
func Write(w http.ResponseWriter, r *http.Request, payload interface{}, status int) {
	op := errors.Op("payload.Write")

	encoded, err := json.Marshal(payload)
	if err != nil {
		handleInternalServerError(w, r, errors.E(op, err))
		return
	}

	// Set the appropriate headers to indicate gzip compression and JSON content type
	w.Header().Add("Content-Encoding", "gzip")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	if err != nil {
		panic(errors.E(op, err))
	}

	defer gz.Close()

	if _, err = gz.Write(encoded); err != nil {
		panic(errors.E(op, err))
	}
}

// handleInternalServerError writes the given error to stderr and returns a
// 500 response with a default message.
func handleInternalServerError(w http.ResponseWriter, _ *http.Request, e error) {
	// log.AlarmWithContext(r.Context(), e)
	log.Printf("Err: %v", e)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if _, err := w.Write(encodedErrResp); err != nil {
		panic(errors.E(errors.Op("payload.handleInternalServerError"), err))
	}
}
