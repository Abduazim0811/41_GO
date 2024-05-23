package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"Homework_41/models"
)

func main() {
	http.HandleFunc("/greet/", GreetHandler)

	fmt.Println("Microservice listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}


func GreetHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}
	name := pathSegments[2]

	url := fmt.Sprintf("http://localhost:8080/hello?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch greeting", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var greeting models.GreetingResponse
	if err := json.NewDecoder(resp.Body).Decode(&greeting); err != nil {
		http.Error(w, "Failed to decode greeting response", http.StatusInternalServerError)
		return
	}

	time.Sleep(500 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(greeting); err != nil {
		http.Error(w, "Failed to encode greeting response", http.StatusInternalServerError)
		return
	}
}
