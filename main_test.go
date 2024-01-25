package main

import (
	"bytes"
	"estiam/dictionary"
	"estiam/middleware"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestIntegration(t *testing.T) {
	// Initialize the dictionary with BoltDB
	dict, err := dictionary.New()
	if err != nil {
		t.Fatalf("Error initializing dictionary: %v", err)
	}
	defer dict.Close()

	// Create a new router
	r := mux.NewRouter()

	// Add the authentication middleware to the router
	r.Use(middleware.AuthenticationMiddleware)

	// Add the logging middleware to the router
	r.Use(middleware.LoggingMiddleware)

	// Define routes
	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		actionAdd(dict, w, r)
	}).Methods("POST")

	r.HandleFunc("/define/{word}", func(w http.ResponseWriter, r *http.Request) {
		actionDefine(dict, w, r)
	}).Methods("GET")

	r.HandleFunc("/remove/{word}", func(w http.ResponseWriter, r *http.Request) {
		actionRemove(dict, w, r)
	}).Methods("DELETE")

	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		actionList(dict, w, r)
	}).Methods("GET")

	// Create a test server
	testServer := httptest.NewServer(r)
	defer testServer.Close()

	// Integration test for the "/add" route
	t.Run("IntegrationTest_AddRoute", func(t *testing.T) {
		url := fmt.Sprintf("%s/add", testServer.URL)
		payload := []byte(`{"word": "testword", "definition": "testdefinition"}`)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Set a timeout for the HTTP client
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error sending request: %v", err)
		}
		defer resp.Body.Close()

		// Check if the response status code is OK
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
	})
}
