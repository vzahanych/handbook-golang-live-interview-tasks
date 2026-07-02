package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "sort"
)

// We want to build an API endpoint to analyze web server logs.
//
// 1.  **Implement the `analyzeHandler` function.**
//     - It should decode the incoming JSON request body into the `LogRequest` struct.
//     - Handle potential JSON decoding errors gracefully.
//
// 2.  **Process the logs to find the top 3 most visited pages.**
//     - Count the occurrences of each `path`.
//     - Determine the top 3 paths based on their counts.
//
// 3.  **Send the response.**
//     - The response should be a JSON object matching the `TopPagesResponse` struct.
//     - The `top_pages` slice should be sorted in descending order of `count`.
//     - Set the correct `Content-Type` header to `application/json`.
//
// Request Example:
// {
//   "logs": [
//     {"path": "/home", "status_code": 200},
//     {"path": "/products", "status_code": 200},
//     {"path": "/home", "status_code": 200},
//     {"path": "/about", "status_code": 200},
//     {"path": "/products", "status_code": 200},
//     {"path": "/home", "status_code": 500}
//   ]
// }
//
// Response Example:
// {
//   "top_pages": [
//     {"path": "/home", "count": 3},
//     {"path": "/products", "count": 2},
//     {"path": "/about", "count": 1}
//   ]
// }

type LogRequest struct {
	Logs []LogEntry `json:"logs"`
}

type LogEntry struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
}

type TopPagesResponse struct {
	TopPages []PageCount `json:"top_pages"`
}

type PageCount struct {
	Path  string `json:"path"`
	Count int    `json:"count"`
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement your logic here.

	// A placeholder response to show it's working.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "not implemented yet"})
}

func main() {
	http.HandleFunc("/analyze/top-pages", analyzeHandler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
