package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Define the Activity struct
type Activity struct {
	StudentID string    `json:"student_id"`
	URL       string    `json:"url"`
	Timestamp time.Time `json:"timestamp"`
}

// In-memory list to store activities
var activities []Activity

func main() {
	http.HandleFunc("/activity", handleActivity)
	http.HandleFunc("/report", handleReport)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleActivity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var activity Activity
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	activities = append(activities, activity)
	w.WriteHeader(http.StatusCreated)
}

func handleReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
