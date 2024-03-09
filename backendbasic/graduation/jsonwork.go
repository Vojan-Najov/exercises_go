package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Request struct {
	Format string `json:"format"` // Format, as time.Format. If empty, use time.RFC333
	TZ     string `json:"tz"`     // TZ, as time.LoadLocation. If empty, use time.Local
}

// The time, formatted according to the request's Format and TZ.
// no need for omitempty here; we'll never send a zero time.
type Response struct {
	Time string `json:"time"`
}

// no need for omitempty here; we'll never send a zero time.
type Error struct {
	Error string `json:"error"`
}

// http handler: writes current time as JSON object (`{"Time": <time>}`)
func getTime(w http.ResponseWriter, r *http.Request) {
	var req Request
	w.Header().Set("Content-Type", "encoding/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(400) // bad request
		json.NewEncoder(w).Encode(Error{err.Error()})
		return
	}
	r.Body.Close() // always close request bodies when you're done with them.
	var tz *time.Location = time.Local
	if req.TZ != "" {
		var err error
		tz, err = time.LoadLocation(req.TZ)
		if err != nil || tz == nil {
			w.WriteHeader(400) // bad request
			json.NewEncoder(w).Encode(Error{err.Error()})
			return
		}
	}
	format := time.RFC3339
	if req.Format != "" {
		format = req.Format
	}

	resp := Response{time.Now().In(tz).Format(format)}
	json.NewEncoder(w).Encode(resp)
}

// Helpful generic functions
// Reading and writing JSON can seem tedious. The following generic functions can help
// reduce boilerplate and help you avoid common 'gotchas', like forgetting to close
// the response body.

// ReadJSON reads a JSON object from an io.ReadCloser, closing the reader when it's
// done. It's primarily useful for reading JSON from *http.Request.Body.
func ReadJSON[T any](r io.ReadCloser) (T, error) {
	var v T                               // declare a variable of type T
	err := json.NewDecoder(r).Decode(&v)  // decode the JSON into v
	return v, errors.Join(err, r.Close()) // close the reader and return any errors.
}

// WriteJSON writes a JSON object to a http.ResponseWriter,
// setting the Content-Type header to application/json.
func WriteJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// Similarly, you may wish to define some helper functions for your own JSON APIs.

// WriteError logs an error, then writes it as a JSON object in the form
// {"error": <error>}, setting the Content-Type header to application/json.
func WriteError(w http.ResponseWriter, err error, code int) {
	// log the error; http.StatusText gets "Not Found" from 4040, etc.
	log.Printf("%d %v: %v", code, http.StatusText(code), err)
	w.Header().Set("Content-Type", "encoding/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Error{err.Error()})
}

// The combination of anonymous structs and generics allows us to write much
// more compact handlers without dragging in a full-blown web framework
// Let's rewrite the logic of getTime to use this technique.

// http handler: writes current time as JSON object (`{"Time": <time>}`)
func getTimeCompact(w http.ResponseWriter, r *http.Request) {
	req, err := ReadJSON[struct{ TZ, Format string }](r.Body)
	if err != nil {
		WriteError(w, err, 400)
		return
	}
	var tz *time.Location = time.Local
	if req.TZ != "" {
		var err error
		tz, err = time.LoadLocation(req.TZ)
		if err != nil {
			WriteError(w, err, 400)
			return
		}
	}
	format := time.RFC3339
	if req.Format != "" {
		format = req.Format
	}
	WriteJSON(w, Response{time.Now().In(tz).Format(format)})
}
