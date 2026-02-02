package ratelimiter

import "time"

type Bucket struct {
	tokens     float64
	capacity   float64
	refillRate float64
	lastRefill time.Time
}

func NewBucket(capacity, refillRate float64) *Bucket {
	return &Bucket{
		tokens:     capacity,
		capacity:   capacity,
		refillRate: capacity,
		lastRefill: time.Now(),
	}
}

func (b *Bucket) refill(now time.Time) {
	elapsed := now.Sub(b.lastRefill).Seconds()

	if elapsed <= 0 {
		return
	}

	b.tokens += elapsed * b.refillRate

	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}

	b.lastRefill = now
}

func (b *Bucket) allow(now time.Time) bool {
	b.refill(now)

	if b.tokens >= 1 {
		b.tokens -= 1
		return true
	}
	return false
}
