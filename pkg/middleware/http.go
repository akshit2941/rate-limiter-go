package middleware

import (
	"encoding/json"
	"net/http"
	"rate-limiter-go/internal/ratelimiter"
	"strconv"
)

func RateLimitMiddleware(limiter ratelimiter.Limiter) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				apiKey = "anonymous"
			}

			result := limiter.Allow(apiKey)

			w.Header().Set("X-RateLimit-Limit", strconv.Itoa(result.Limit))
			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(result.Remaining))

			if !result.Allowed {
				w.Header().Set(
					"Retry-After",
					strconv.Itoa(int(result.RetryAfter.Seconds())),
				)

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
