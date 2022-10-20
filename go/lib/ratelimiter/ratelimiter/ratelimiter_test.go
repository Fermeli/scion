package ratelimiter

import (
	"testing"
	"time"
)

func TestApplyReturnsFalseIfNoRateLimitsHaveBeenAdded(t *testing.T) {

	rateLimiter := NewRateLimiter()

	now := time.Now()
	tests := []struct {
		identifier string
		rate       time.Duration
		cbs        int64
		pktLen     int64
		now        time.Time
		want       bool
	}{
		{"nrored", 1, 1, 1, now, false},
		{"fjriojvoiq", 1, 2, 1, now.Add(10), false},
		{"joqvkfd", 50, 10, 1, now.Add(100), false},
		{"124855", 30, 40, 1, now.Add(1000), false},
	}

	for _, test := range tests {
		res := rateLimiter.Apply(test.identifier, test.pktLen, test.now)
		if res != test.want {
			t.Errorf("Error")
		}
	}
}

func TestApplyReturnsFalseIfIdentfierHaveNotBeenAddedAndTrueOtherwise(t *testing.T) {

	rateLimiter := NewRateLimiter()
	now := time.Now()

	rateLimiter.AddRatelimit("identifier", 1000000000, 5, now)

	tests := []struct {
		identifier string
		rate       time.Duration
		cbs        int64
		pktLen     int64
		now        time.Time
		want       bool
	}{
		{"identifier", 1, 1, 1, now.Add(10), true},
		{"identifier1", 1, 2, 1, now.Add(10), false},
		{"ientifier", 50, 10, 1, now.Add(100), false},
		{"124855fezfzezffze", 30, 40, 1, now.Add(1000), false},
	}

	for _, test := range tests {
		res := rateLimiter.Apply(test.identifier, test.pktLen, test.now)
		if res != test.want {
			t.Errorf("Error")
		}
	}
}

func TestGetBurstSizeReturnsAnErrorWithUnknownIdentifier(t *testing.T) {

	rateLimiter := NewRateLimiter()

	tests := []struct {
		identifier string
		want       int64
	}{
		{"nrored", -1},
		{"fjriojvoiq", -1},
		{"joqvkfd", -1},
		{"124855", -1},
	}

	for _, test := range tests {
		res, err := rateLimiter.GetBurstSize(test.identifier)
		if res != test.want || err == nil {
			t.Errorf("Error")
		}
	}
}

func TestGetRateReturnsAnErrorWithUnknownIdentifier(t *testing.T) {

	rateLimiter := NewRateLimiter()

	tests := []struct {
		identifier string
		want       float64
	}{
		{"nrored", -1},
		{"fjriojvoiq", -1},
		{"joqvkfd", -1},
		{"124855", -1},
	}

	for _, test := range tests {
		res, err := rateLimiter.GetRate(test.identifier)
		if res != test.want || err == nil {
			t.Errorf("Error")
		}
	}
}

func TestSetBurstSizeReturnsAnErrorWithUnknownIdentifier(t *testing.T) {

	rateLimiter := NewRateLimiter()

	tests := []struct {
		identifier string
	}{
		{"nrored"},
		{"fjriojvoiq"},
		{"joqvkfd"},
		{"124855"},
	}

	for _, test := range tests {
		err := rateLimiter.SetBurstSize(test.identifier, 5)
		if err == nil {
			t.Errorf("Error")
		}
	}
}

func TestSetRateReturnsAnErrorWithUnknownIdentifier(t *testing.T) {

	rateLimiter := NewRateLimiter()

	tests := []struct {
		identifier string
	}{
		{"nrored"},
		{"fjriojvoiq"},
		{"joqvkfd"},
		{"124855"},
	}

	for _, test := range tests {
		err := rateLimiter.SetRate(test.identifier, 5)
		if err == nil {
			t.Errorf("Error")
		}
	}
}
