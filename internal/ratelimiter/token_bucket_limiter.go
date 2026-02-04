package ratelimiter

import "time"

type TokenBucketLimiter struct {
	store Store
}

func NewTokenBucketLimiter(store Store) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		store: store,
	}
}

func (l *TokenBucketLimiter) Allow(key string) bool {
	bucket := l.store.Get(key)

	now := time.Now()

	return bucket.allow(now)
}
