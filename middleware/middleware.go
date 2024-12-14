package middleware

import (
	"net/http"
)

func RequireRole(requiredRole string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the role from headers (or any other authentication system you're using)
		role := r.Header.Get("Role")
		if role != requiredRole {
			http.Error(w, "Forbidden: Insufficient role permissions", http.StatusForbidden)
			return
		}
		// Call the next handler if the role matches
		next(w, r)
	}
}
