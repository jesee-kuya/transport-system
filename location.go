package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	currentLocation = Location{Latitude: -0.066285, Longitude: 34.774352}
	mu              sync.Mutex
)

func UpdateLocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var location Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	currentLocation = location
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		mu.Lock()
		location := currentLocation
		mu.Unlock()

		data, _ := json.Marshal(location)
		fmt.Fprintf(w, "data: %s\n\n", data)
		w.(http.Flusher).Flush()

		time.Sleep(2 * time.Second) // Update every 2 seconds
	}
}
