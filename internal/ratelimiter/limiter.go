package raetlimiter

type Limiter interface {
	Allow(key string) bool
}
