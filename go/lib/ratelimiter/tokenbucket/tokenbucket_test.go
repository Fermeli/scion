package tokenbucket

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestNoPacketCanBeSentAtInitTime(t *testing.T) {

	now := time.Now()

	tests := []struct {
		token  TokenBucket
		tokens int64
		now    time.Time
		want   bool
	}{
		{NewTokenBucket(now.Add(10), 100, 2), 15, now.Add(10), false},
		{NewTokenBucket(now.Add(10), 100, 2), 15, now.Add(10), false},
		{NewTokenBucket(now.Add(100), 1000, 3), 20, now.Add(100), false},
		{NewTokenBucket(now.Add(1000), 10000, 4), 50, now.Add(1000), false},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.want {
			fmt.Println(test.token.currentTokens)
			t.Errorf("Error")
		}
	}

}

func TestBucketIsUpdatedProperlyAfterSuccess(t *testing.T) {

	now := time.Now()
	tests := []struct {
		token      TokenBucket
		tokens     int64
		now        time.Time
		wantReturn bool
		wantTokens int64
	}{
		{NewTokenBucket(now, 10, 10000000000), 10, now.Add(1), true, 0},
		{NewTokenBucket(now, 100, 100000000000), 15, now.Add(1), true, 85},
		{NewTokenBucket(now, 1000, 1000000000000), 20, now.Add(1), true, 980},
		{NewTokenBucket(now, 10000, 10000000000000), 50, now.Add(1), true, 9950},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.wantReturn || test.token.currentTokens != test.wantTokens {
			t.Errorf("Error")
		}
	}
}

func TestBucketIsEmptyForEveybodyAfterUsingAllTokensAtOnce(t *testing.T) {

	now := time.Now()
	tests := []struct {
		token      TokenBucket
		tokens     int64
		now        time.Time
		wantReturn bool
		wantTokens int64
	}{
		{NewTokenBucket(now, 10, 10000000000), 10, now.Add(1), true, 0},
		{NewTokenBucket(now, 100, 100000000000), 100, now.Add(10), true, 0},
		{NewTokenBucket(now, 1000, 50000000000), 1000, now.Add(20), true, 0},
		{NewTokenBucket(now, 10000, 250000000000), 10000, now.Add(40), true, 0},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.wantReturn || test.token.currentTokens != test.wantTokens {
			t.Errorf("Error")
		}
	}

}

func TestLargePacketsSmallRates(t *testing.T) {

	now := time.Now()

	tests := []struct {
		token  TokenBucket
		tokens int64
		now    time.Time
		want   bool
	}{
		{NewTokenBucket(now, 1, 100), 1000000, now, false},
		{NewTokenBucket(now.Add(10), 1, 100), 55555555, now.Add(10), false},
		{NewTokenBucket(now.Add(100), 1, 100), 99999999999, now.Add(100), false},
		{NewTokenBucket(now.Add(1000), 1, 100), math.MaxInt64, now.Add(1000), false},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.want {
			t.Errorf("Error")
		}
	}
}

func TestBurstyTraffic(t *testing.T) {

	now := time.Now()
	token := NewTokenBucket(now, 100000, 2000000000000)
	tests := []struct {
		tokens int64
		now    time.Time
		want   bool
	}{
		{1, now.Add(51), true},
		{1, now.Add(52), true},
		{1, now.Add(53), true},
		{1, now.Add(54), true},
		{1, now.Add(55), true},
		{50000, now.Add(56), true},
		{50000, now.Add(57), true},
		{50000, now.Add(58), false},
		{50000, now.Add(59), false},
		{50000, now.Add(60), false},
	}

	for _, test := range tests {
		res := token.Apply(test.tokens, test.now)
		if res != test.want {
			t.Errorf("Error")
		}
	}
}

func TestStartSendingAgainAfterALongTime(t *testing.T) {

	now := time.Now()
	token := NewTokenBucket(now, 100000, 2000000000000)
	tests := []struct {
		tokens int64
		now    time.Time
		want   bool
	}{
		{1, now.Add(51), true},
		{1, now.Add(52), true},
		{1, now.Add(53), true},
		{1, now.Add(54), true},
		{1, now.Add(55), true},
		{50000, now.Add(555556), true},
		{50000, now.Add(555557), true},
		{50000, now.Add(555558), false},
		{50000, now.Add(555559), false},
		{50000, now.Add(555560), false},
	}

	for _, test := range tests {
		res := token.Apply(test.tokens, test.now)
		if res != test.want {
			t.Errorf("Error")
		}
	}
}

func TestSendOverALongPeriod(t *testing.T) {

	type testStruct struct {
		tokens int64
		now    time.Time
		want   bool
	}

	now := time.Now()
	token := NewTokenBucket(now, 50, 1000000000)
	var tests []testStruct
	var i time.Duration

	for i = 1; i < 50; i += 1 {
		tests = append(tests, testStruct{int64(1 + (i%2)*60), now.Add(i), i%2 == 0})
	}

	for _, test := range tests {
		res := token.Apply(test.tokens, test.now)
		if res != test.want {
			t.Errorf("Error")
		}
	}
}

func TestLargeRates(t *testing.T) {

	now := time.Now()

	tests := []struct {
		token  TokenBucket
		tokens int64
		now    time.Time
		want   bool
	}{
		{NewTokenBucket(now, 10, 1000000), 1, now, false},
		{NewTokenBucket(now.Add(10), 10, 55555555), 10, now.Add(10), false},
		{NewTokenBucket(now.Add(100), 10, 99999999999), 100, now.Add(100), false},
		{NewTokenBucket(now, 99999999999, 99999999999), 1, now.Add(9999999999), true},
		{NewTokenBucket(now.Add(1000), 10, math.MaxInt64), math.MaxInt64, now.Add(1000), false},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.want {
			t.Errorf("Error")
		}
	}
}

func TestPacketOfSize1AfterLongTimeIsASuccess(t *testing.T) {

	now := time.Now()

	tests := []struct {
		token  TokenBucket
		tokens int64
		now    time.Time
		want   bool
	}{
		{NewTokenBucket(now.Add(10), 45656, 2), 1, now.Add(10000000000), true},
		{NewTokenBucket(now.Add(10), 8944856, 2), 1, now.Add(10000000000), true},
		{NewTokenBucket(now.Add(100), 48968465, 3), 1, now.Add(1000000000000), true},
		{NewTokenBucket(now.Add(1000), 84654878465, 4), 1, now.Add(10000000000000), true},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.want {
			fmt.Println(test.token.currentTokens)
			t.Errorf("Error")
		}
	}

}

func TestTheBucketShouldNotBeFilledWhenTheRateIsZero(t *testing.T) {

	now := time.Now()

	tests := []struct {
		token  TokenBucket
		tokens int64
		now    time.Time
		want   bool
	}{
		{NewTokenBucket(now.Add(10), 847646, 0), 1, now.Add(10000), false},
		{NewTokenBucket(now.Add(10), 87464865, 0), 1, now.Add(10000), false},
		{NewTokenBucket(now.Add(100), 7894384965, 0), 1, now.Add(100000), false},
		{NewTokenBucket(now.Add(1000), 96586586, 0), 1, now.Add(10000000), false},
	}

	for _, test := range tests {
		if test.token.Apply(test.tokens, test.now) != test.want {
			fmt.Println(test.token.currentTokens)
			t.Errorf("Error")
		}
	}

}
