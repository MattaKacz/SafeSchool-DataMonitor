package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// Define the Activity struct
type Activity struct {
	StudentID string    `json:"student_id"`
	URL       string    `json:"url"`
	Timestamp time.Time `json:"timestamp"`
}

// In-memory list to store activities
var activities []Activity

var db *sql.DB

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("/activity", handleActivity)
    http.HandleFunc("/report", handleReport)

    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


func initDB() {
    connStr := "user=safeschool_user password=yourpassword dbname=safeschool sslmode=disable"
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database!")
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

    query := `INSERT INTO activities (student_id, url, timestamp) VALUES ($1, $2, $3)`
    _, err = db.Exec(query, activity.StudentID, activity.URL, activity.Timestamp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// Store the activity in the database
func handleReport(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT student_id, url, timestamp FROM activities")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var activities []Activity
    for rows.Next() {
        var activity Activity
        err := rows.Scan(&activity.StudentID, &activity.URL, &activity.Timestamp)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        activities = append(activities, activity)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(activities)
}
