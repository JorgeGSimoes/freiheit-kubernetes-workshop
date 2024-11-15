package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"
)

// Response structure to send back the random number
type RandomNumberResponse struct {
	RandomNumber int `json:"random_number"`
}

//const maxConcurrentRequests = 5 // Limit the number of concurrent requests
//var semaphore = make(chan struct{}, maxConcurrentRequests)

// Generates a random number in the specified range [min, max]
func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Handler for the /random endpoint
func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Acquire semaphore
	//semaphore <- struct{}{}
	//defer func() {
	//	// Release semaphore
	//	<-semaphore
	//}()

	// Simulate computation delay
	time.Sleep(500 * time.Millisecond)
	fmt.Println("generating random number")

	// Generate the random number and send response
	randomNumber := generateRandomNumber(1, 100)
	response := RandomNumberResponse{RandomNumber: randomNumber}

	// Set header to application/json and encode the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler to serve the static HTML file
func serveHTMLHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Serve the HTML file
	http.ServeFile(w, r, filepath.Join(".", "index.html"))
}

func main() {
	// Serve the HTML file at the root URL "/"
	http.HandleFunc("/", serveHTMLHandler)

	// Handle /random endpoint
	http.HandleFunc("/random", randomNumberHandler)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
