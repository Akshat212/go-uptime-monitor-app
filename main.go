package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Uptime monitor is running 🚀")
}

func main() {
	// Basic routing
	http.HandleFunc("/health", healthHandler)

	// Start Server
	port := "8080"
	log.Printf("Server is running and listening at port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}