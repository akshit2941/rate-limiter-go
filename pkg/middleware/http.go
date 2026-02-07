package middleware

import (
	"encoding/json"
	"net/http"
	"rate-limiter-go/internal/ratelimiter"
)

func RateLimitMiddleware(limiter ratelimiter.Limiter) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				apiKey = "anonymous"
			}

			allowed := limiter.Allow(apiKey)

			if !allowed {
				w.WriteHeader(http.StatusTooManyRequests)

				response := map[string]string{
					"error": "rate limit exceeded",
				}

				_ = json.NewEncoder(w).Encode(response)

				return
			}
			next.ServeHTTP(w, r)

		})
	}
}
