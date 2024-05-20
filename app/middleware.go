package app

import (
	"net/http"
	"video_annotator/constants"
)

// fetch from secrets
const apiKey = "secret-api-key"

// WithProtectedAuth if possible implement authorization permission and roles
func WithProtectedAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")
		if key != apiKey {
			http.Error(w, "Unauthorized", constants.HttpUnauthorised)
			return
		}
		next.ServeHTTP(w, r)
	}
}
