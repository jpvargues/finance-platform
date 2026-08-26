package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusCode struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s StatusCode
	s.Status = "ok"
	userJSON, err := json.Marshal(s)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}
	w.Write(userJSON)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
