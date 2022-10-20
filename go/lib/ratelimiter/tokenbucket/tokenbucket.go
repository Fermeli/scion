package tokenbucket

import (
	"math"
	"time"
)

// Data structure representing a token bucket
type TokenBucket struct {
	currentTokens   int64
	lastTimeApplied time.Time
	cbs             int64
	//in tokens/seconds
	cir float64
}

// NewTokenBucket initializes the fields of the token bucket and returns it given the rate
// in tokens/seconds, the cbs in tokens and the initial time.
func NewTokenBucket(initialTime time.Time, cbs int64, rate float64) TokenBucket {
	t := TokenBucket{}
	t.cir = rate * math.Pow(10, -9)
	t.cbs = cbs
	t.currentTokens = 0
	t.lastTimeApplied = initialTime
	return t
}

// SetRate sets the value of the rate to the one given
func (t *TokenBucket) SetRate(rate float64) {
	t.cir = rate * math.Pow(10, -9)

	//if new rate is 0 we empty the bucket
	if rate == 0 {
		t.currentTokens = 0
	}

}

// SetBurstSize sets the value of the cbs to the one given
func (t *TokenBucket) SetBurstSize(cbs int64) {
	t.cbs = cbs
}

// GetRate returns the rate in tokens/second
func (t *TokenBucket) GetRate() float64 {
	return t.cir * math.Pow(10, 9)
}

// GetBurstSize returns the cbs (Committed burst size)
func (t *TokenBucket) GetBurstSize() int64 {
	return t.cbs
}

// SetBurstSizeAndRate sets the value of the cbs and of the rate at the ones given
func (t *TokenBucket) SetBurstSizeAndRate(cbs int64, rate float64) {
	t.SetBurstSize(cbs)
	t.SetRate(rate)
}

// GetBurstSizeAndRate returns the cbs and of the rate
func (t *TokenBucket) GetBurstSizeAndRate() (cbs int64, rate float64) {
	return t.GetBurstSize(), t.GetRate()
}

// Apply applies the token bucket algorithm. If there are not enough tokens in the buckets
// for the given pkt_len, the token bucket is not updated. If the token bucket is updated
// succesfully, the function returns true, otherwise it returns false. The token must be
// initialized. The currentTokens and lastTimeApplied attributes of the TokenBucket t are
// updated only if the call to apply is a success.
func (t *TokenBucket) Apply(tokens int64, now time.Time) bool {
	// You cannot update at a time preceding a previous successful call to the function
	if now.Before(t.lastTimeApplied) {
		return false
	}

	// The tokens added between two times t1 t2 as (t2 - t1) * CIR.
	t.currentTokens += int64(math.Floor(float64(now.Sub(t.lastTimeApplied)) * t.cir))

	// There cannot be more than CBS tokens in the bucket.
	if t.currentTokens > t.cbs {
		t.currentTokens = t.cbs
	}

	t.lastTimeApplied = now

	// Check if the number of tokens required is smaller than the number of tokens in the bucket
	if t.currentTokens >= tokens {
		t.currentTokens -= tokens

		return true
	} else {
		return false
	}
}
