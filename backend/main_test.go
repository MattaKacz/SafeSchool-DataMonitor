package main

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandleActivity(t *testing.T) {
    jsonData := []byte(`{"student_id":"123","url":"https://example.com","timestamp":"2025-04-29T14:00:00Z"}`)
    req, err := http.NewRequest("POST", "/activity", bytes.NewBuffer(jsonData))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleActivity)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}

func TestHandleReport(t *testing.T) {
    req, err := http.NewRequest("GET", "/report", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handleReport)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    if rr.Body.Len() == 0 {
        t.Errorf("handler returned empty body")
    }
}
