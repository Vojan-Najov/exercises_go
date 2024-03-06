package main

import (
	"bytes"
	"context"
	"encoding/json"
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

var client = &http.Client{Timeout: 2 * time.Second}

func sendRequest(tz, format string) {
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(Request{TZ: tz, Format: format})
	log.Printf("request body: %v", body)
	req, err := http.NewRequestWithContext(
		context.TODO(), "GET", "http://localhost:8080", body,
	)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Write(os.Stdout)
	resp.Body.Close() // always close response bodies when you're do with them
}

func main() {
	server := http.Server{Addr: ":8080", Handler: http.HandlerFunc(getTime)}
	go server.ListenAndServe()

	sendRequest("", "") // rely on defaults
	sendRequest("America/Los_Angeles", time.RFC822Z)
	sendRequest("America/New_York", time.RFC822Z)
	sendRequest("faketz", "")
}
