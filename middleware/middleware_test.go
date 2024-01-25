package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAuthenticationMiddleware_ValidToken(t *testing.T) {
	// Create a test handler that will be wrapped by the middleware
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a request with a valid token
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "viverk")

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Use the middleware and the test handler
	AuthenticationMiddleware(testHandler).ServeHTTP(recorder, req)

	// Check if the response code is OK (200)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestAuthenticationMiddleware_InvalidToken(t *testing.T) {
	// Create a test handler that will be wrapped by the middleware
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler should not be called with an invalid token")
		w.WriteHeader(http.StatusOK)
	})

	// Create a request with an invalid token
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "invalidToken")

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Use the middleware and the test handler
	AuthenticationMiddleware(testHandler).ServeHTTP(recorder, req)

	// Check if the response code is Unauthorized (401)
	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, recorder.Code)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	// Create a buffer to capture log output
	var logBuffer bytes.Buffer

	// Create a test handler that will be wrapped by the middleware
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a request
	req := httptest.NewRequest("GET", "/list", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Use the middleware and the test handler
	LoggingMiddlewareWithLogger(testHandler, log.New(&logBuffer, "", 0)).ServeHTTP(recorder, req)

	// Check if the log output contains the expected information
	expectedLogOutput := "Method: GET - URI: /list - Protocol: HTTP/1.1"
	if !containsLogOutput(logBuffer.String(), expectedLogOutput) {
		t.Errorf("Log output does not contain the expected information:\nExpected: %s\nActual: %s",
			expectedLogOutput, logBuffer.String())
	}
}

// LoggingMiddlewareWithLogger is the modified middleware to use a custom logger.
func LoggingMiddlewareWithLogger(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time of the request
		start := time.Now()

		// Create a buffer to capture the response
		recorder := httptest.NewRecorder()

		// Serve the request to the next middleware or handler in the chain
		next.ServeHTTP(recorder, r)

		// Record the end time of the request
		end := time.Now()

		// Calculate the duration of the request
		duration := end.Sub(start)

		// Log the request information
		logger.Printf("Method: %s - URI: %s - Protocol: %s - Duration: %v",
			r.Method, r.RequestURI, r.Proto, duration)

		// Copy the recorded response to the original response writer
		for k, v := range recorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(recorder.Code)
		recorder.Body.WriteTo(w)
	})
}

// containsLogOutput checks if the log output contains the expected information.
func containsLogOutput(logOutput, expectedInfo string) bool {
	return strings.Contains(logOutput, expectedInfo)
}
