package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	currentLocation = Location{Latitude: -1.286389, Longitude: 36.817223} // Default location (Nairobi)
	mu              sync.Mutex
)

func updateLocationHandler(w http.ResponseWriter, r *http.Request) {
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

func streamLocationHandler(w http.ResponseWriter, r *http.Request) {
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

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html") // Serve the index.html file
}

func main() {
	http.HandleFunc("/", serveIndex) // Serve index.html at root
	http.HandleFunc("/update-location", updateLocationHandler)
	http.HandleFunc("/stream-location", streamLocationHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
