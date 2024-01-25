package middleware

import (
	"net/http"
)

// AuthenticationMiddleware checks for the presence and validity of an authentication token.
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the authentication token from the request header
		token := r.Header.Get("Authorization")

		// Check if the token is valid (you should implement your own logic here)
		if isValidToken(token) {
			// If valid, proceed to the next middleware or handler in the chain
			next.ServeHTTP(w, r)
		} else {
			// If not valid, return an unauthorized response
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

// isValidToken is a placeholder function that you should replace with your own token validation logic.
func isValidToken(token string) bool {
	// Implement your token validation logic here
	// Example: Check if the token is present and matches a predefined value
	return token == "viverk"
}
