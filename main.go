package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Akshat212/go-uptime-monitor-app/internal/monitor"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Uptime monitor is running 🚀")
}

func addURLHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	monitor.AddURL(req.URL)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Monitoring %s\n", req.URL)
}

func listURLsHandler(w http.ResponseWriter, r *http.Request) {
	urls := monitor.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}

func main() {
	// Basic routing
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/monitor", addURLHandler)
	http.HandleFunc("/monitored", listURLsHandler)

	// Start Server
	port := "8080"
	log.Printf("Server is running and listening at port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}