package ratelimiter

import (
	"fmt"
	"time"

	"github.com/scionproto/scion/go/lib/ratelimiter/tokenbucket"
)

type RateLimiter struct {
	// The 10 bytes array is derived from the SrcIA of type addr.IA and the EgressID
	// of type uint16. The SrcIA is 8 bytes long and the EgressID is 2 bytes long.
	// The identifier of the rate limiter is built from the concatenation of these
	// two values so it is 10 bytes long.
	buckets map[[10]byte]*tokenbucket.TokenBucket
}

// NewRateLimiter initalizes and returns a RateLimiter. The map of buckets is empty at
// initialization.
func NewRateLimiter() RateLimiter {
	r := RateLimiter{}
	r.Clear()
	return r
}

// Clear clears the buckets of the RateLimiter
func (r *RateLimiter) Clear() {
	r.buckets = make(map[[10]byte]*tokenbucket.TokenBucket)
}

// AddRatelimit adds a new identifier and initializes the fields of its bucket
func (r *RateLimiter) AddRatelimit(identifier [10]byte, rate float64, cbs int64, now time.Time) {
	newTokenBucket := tokenbucket.NewTokenBucket(now, cbs, rate)
	r.buckets[identifier] = &newTokenBucket
}

// SetRate sets the value of the time interval of the token bucket of the given identifier to the
// one given
func (r *RateLimiter) SetRate(identifier [10]byte, rate float64) error {
	tokenBucket, keyPresent := r.buckets[identifier]

	if keyPresent {
		tokenBucket.SetRate(rate)
		return nil
	}

	return fmt.Errorf("Identifier '%s' have not been added yet", identifier)
}

// SetBurstSize sets the value of the cbs of the token bucket of the given identifier to
// the one given
func (r *RateLimiter) SetBurstSize(identifier [10]byte, cbs int64) error {
	tokenBucket, keyPresent := r.buckets[identifier]

	if keyPresent {
		tokenBucket.SetBurstSize(cbs)
		return nil
	}

	return fmt.Errorf("Identifier '%s' have not been added yet", identifier)
}

// GetRate returns T the time interval of the token bucket of the given identifier, or -1 and an
// error if the identifier is not present in the map
func (r *RateLimiter) GetRate(identifier [10]byte) (float64, error) {
	tokenBucket, keyPresent := r.buckets[identifier]

	if !keyPresent {
		return -1, fmt.Errorf("Identifier '%s' have not been added yet", identifier)
	}

	return tokenBucket.GetRate(), nil
}

// GetBurstSize returns the cbs (Committed burst size) of the token bucket of the given identifier,
// or -1 and an error if the identifier is not present in the map
func (r *RateLimiter) GetBurstSize(identifier [10]byte) (int64, error) {
	tokenBucket, keyPresent := r.buckets[identifier]

	if !keyPresent {
		return -1, fmt.Errorf("Identifier '%s' have not been added yet", identifier)
	}

	return tokenBucket.GetBurstSize(), nil
}

// GetBurstSizeAndRate returns the cbs and of the rate
func (r *RateLimiter) GetBurstSizeAndRate(identifier [10]byte) (int64, float64, error) {
	tokenBucket, keyPresent := r.buckets[identifier]

	if !keyPresent {
		return -1, -1, fmt.Errorf("Identifier '%s' have not been added yet", identifier)
	}

	return tokenBucket.GetBurstSize(), tokenBucket.GetRate(), nil
}

// SetBurstSizeAndRate sets the value of the cbs and of the rate at the ones given
func (r *RateLimiter) SetBurstSizeAndRate(identifier [10]byte, cbs int64, rate float64) error {
	tokenBucket, keyPresent := r.buckets[identifier]

	if keyPresent {
		tokenBucket.SetBurstSize(cbs)
		tokenBucket.SetRate(rate)
		return nil
	}

	return fmt.Errorf("Identifier '%s' have not been added yet", identifier)
}

// Apply applies the token bucket algorithm to the token bucket of the given identifier.
// Check the doc of TokenBucket.Apply
func (r *RateLimiter) Apply(identifier [10]byte, pktLen int64, now time.Time) bool {
	tokenBucket, keyPresent := r.buckets[identifier]

	if !keyPresent {
		return false
	}
	now = time.Now()

	return tokenBucket.Apply(pktLen, now)
}

// Contains return true if the identifier is contained in the map and false otherwise.
func (r *RateLimiter) Contains(identifier [10]byte) bool {
	_, keyPresent := r.buckets[identifier]
	return keyPresent
}
