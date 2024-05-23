package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"Homework_41/models"
)

func main() {
	http.HandleFunc("/hello", HelloHandler)

	fmt.Println("Microservice listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	length := len(name)

	greeting := models.GreetingResponse{
		Name:   name,
		Length: length,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(greeting); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
