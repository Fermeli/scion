package bloomfilter

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestHasBeenSeenReturnsTrueOnlyForKnownIdentifiersWithAllHashFunctions(t *testing.T) {

	bloomFilter := NewBloomFilter(50000000000, 50)

	tests := []struct {
		identifier string
		want1      bool
		want2      bool
	}{
		{"a", false, true},
		{"ab", false, true},
		{"abc", false, true},
		{"abcd", false, true},
		{"abcde", false, true},
		{"abcdef", false, false},
		{"abcdefg", false, false},
		{"abcdefgh", false, false},
		{"abcdefghi", false, false},
		{"abcdefghij", false, false},
	}

	for i, test := range tests {

		if i < 5 {

			res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), true)

			if err != nil {
				t.Error(err)
			}

			if res != test.want1 {
				fmt.Println(res)
				t.Errorf("error")
			}

		}
	}

	for i, test := range tests {
		res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), true)

		if err != nil {
			t.Error(err)
		}

		if res != test.want2 {
			fmt.Println(i)
			t.Errorf("error")
		}

	}

}

func TesstSlidingBloomFilterHasBeenSeenResetsProperly(t *testing.T) {
	now := time.Now()
	timeStampValidity := time.Duration(50)
	bloomFilter := NewSlidingBloomFilter(50000000, 5, 3, timeStampValidity, now)

	tests := []struct {
		identifier      string
		firstTimeCalled time.Time
		want1           bool
		want2           bool
	}{
		{"abcde", now, false, false},
		{"a", now.Add(5), false, false},
		{"ab", now.Add(10), false, false},
		{"abc", now.Add(15), false, false},
		{"abcd", now.Add(20), false, false},
		{"abcdef", now.Add(50), false, true},
		{"abcdefg", now.Add(55), false, true},
		{"abcdefgh", now.Add(60), false, true},
		{"abcdefghi", now.Add(65), false, true},
		{"abcdefghij", now.Add(70), false, true},
	}

	for _, test := range tests {

		res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), test.firstTimeCalled)

		if err != nil {
			t.Error(err)
		}

		if res != test.want1 {
			//fmt.Println(res)
			t.Errorf("error")
		}

	}

	for _, test := range tests {
		res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), now.Add(80))

		if err != nil {
			t.Error(err)
		}

		if res != test.want2 {
			t.Errorf("error")
		}

	}

}

func TestOptimalSlidingBloomFilterHasBeenSeenResetsProperly(t *testing.T) {
	now := time.Now()
	timeStampValidity := 4 * time.Second
	falsePositiveRate := 0.01
	bloomFilter := NewOptimalSlidingBloomFilter(falsePositiveRate, 2*math.Pow10(6), timeStampValidity, now, 3)

	tests := []struct {
		identifier      string
		firstTimeCalled time.Time
		want1           bool
		want2           bool
	}{
		{"abcde", now, false, false},
		{"a", now.Add(5), false, false},
		{"ab", now.Add(10), false, false},
		{"abc", now.Add(15), false, false},
		{"abcd", now.Add(time.Second), false, false},
		{"abcdef", now.Add(2*time.Second + 50), false, true},
		{"abcdefg", now.Add(2*time.Second + 55), false, true},
		{"abcdefgh", now.Add(2*time.Second + 60), false, true},
		{"abcdefghi", now.Add(2*time.Second + 65), false, true},
		{"abcdefghij", now.Add(2*time.Second + 70), false, true},
		{"abcdefghijk", now.Add(4*time.Second + 50), false, true},
		{"abcdefghijkl", now.Add(4*time.Second + 55), false, true},
		{"abcdefghijklm", now.Add(4*time.Second + 60), false, true},
		{"abcdefghijklmn", now.Add(4*time.Second + 70), false, true},
		{"abcdefghijklmno", now.Add(6*time.Second + 50), false, true},
		{"abcdefghijklmnop", now.Add(6*time.Second + 55), false, true},
		{"abcdefghijklmnopq", now.Add(6*time.Second + 60), false, true},
		{"abcdefghijklmnopqr", now.Add(6*time.Second + 65), false, true},
		{"abcdefghijklmnopqrs", now.Add(6*time.Second + 70), false, true},
	}

	for _, test := range tests {

		res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), test.firstTimeCalled)

		if err != nil {
			t.Error(err)
		}

		if res != test.want1 {
			t.Errorf("error")
		}

	}

	for _, test := range tests {
		res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), now.Add(6*time.Second+50))

		if err != nil {
			t.Error(err)
		}

		if res != test.want2 {
			t.Errorf("error")
		}

	}

}

func TestHasBeenSeenReturnsFalseIfTheElementHasNotBeenSeen(t *testing.T) {
	now := time.Now()
	timeStampValidity := time.Second
	falsePositiveRate := 0.001
	bloomFilter := NewOptimalSlidingBloomFilter(falsePositiveRate, math.Pow10(3), timeStampValidity, now, 20)

	tests := []struct {
		identifier      string
		firstTimeCalled time.Time
		want1           bool
	}{
		{"abcde", now, false},
		{"a", now.Add(5), false},
		{"ab", now.Add(10), false},
		{"abc", now.Add(15), false},
		{"abcd", now.Add(time.Second), false},
		{"abcdef", now.Add(2*time.Second + 50), false},
		{"abcdefg", now.Add(2*time.Second + 55), false},
		{"abcdefgh", now.Add(2*time.Second + 60), false},
		{"abcdefghi", now.Add(2*time.Second + 65), false},
		{"abcdefghij", now.Add(2*time.Second + 70), false},
		{"abcdefghijk", now.Add(4*time.Second + 50), false},
		{"abcdefghijkl", now.Add(4*time.Second + 55), false},
		{"abcdefghijklm", now.Add(4*time.Second + 60), false},
		{"abcdefghijklmn", now.Add(4*time.Second + 70), false},
		{"abcdefghijklmno", now.Add(6*time.Second + 50), false},
		{"abcdefghijklmnop", now.Add(6*time.Second + 55), false},
		{"abcdefghijklmnopq", now.Add(6*time.Second + 60), false},
		{"abcdefghijklmnopqr", now.Add(6*time.Second + 65), false},
		{"abcdefghijklmnopqrs", now.Add(6*time.Second + 70), false},
	}

	for _, test := range tests {

		res, err := bloomFilter.HasBeenSeen([]byte(test.identifier), test.firstTimeCalled)

		if err != nil {
			t.Error(err)
		}

		if res != test.want1 {
			t.Errorf("error")
		}

	}

}

func TestHasBeenSeenReturnsTrueIfItIsCalledTwoTimesWithTheSameTimeArgument(t *testing.T) {
	now := time.Now()
	timeStampValidity := 500000 * time.Millisecond
	falsePositiveRate := 0.001
	bloomFilter := NewOptimalSlidingBloomFilter(falsePositiveRate, math.Pow10(3), timeStampValidity, now, 10)

	tests := []struct {
		identifier      string
		firstTimeCalled time.Time
		want1           bool
		want2           bool
	}{
		{"abcde", now, false, true},
		{"a", now.Add(5), false, true},
		{"ab", now.Add(10), false, true},
		{"abc", now.Add(15), false, true},
		{"abcd", now.Add(time.Second), false, true},
		{"abcdef", now.Add(2*time.Second + 50), false, true},
		{"abcdefg", now.Add(2*time.Second + 55), false, true},
		{"abcdefgh", now.Add(2*time.Second + 60), false, true},
		{"abcdefghi", now.Add(2*time.Second + 65), false, true},
		{"abcdefghij", now.Add(2*time.Second + 70), false, true},
		{"abcdefghijk", now.Add(4*time.Second + 50), false, true},
		{"abcdefghijkl", now.Add(4*time.Second + 55), false, true},
		{"abcdefghijklm", now.Add(4*time.Second + 60), false, true},
		{"abcdefghijklmn", now.Add(4*time.Second + 70), false, true},
		{"abcdefghijklmno", now.Add(6*time.Second + 50), false, true},
		{"abcdefghijklmnop", now.Add(6*time.Second + 55), false, true},
		{"abcdefghijklmnopq", now.Add(6*time.Second + 60), false, true},
		{"abcdefghijklmnopqr", now.Add(6*time.Second + 65), false, true},
		{"abcdefghijklmnopqrs", now.Add(6*time.Second + 70), false, true},
	}

	for _, test := range tests {

		res1, err := bloomFilter.HasBeenSeen([]byte(test.identifier), test.firstTimeCalled)

		if err != nil {
			t.Error(err)
		}

		if res1 != test.want1 {
			t.Errorf("error")
		}

		res2, err := bloomFilter.HasBeenSeen([]byte(test.identifier), test.firstTimeCalled)

		if err != nil {
			t.Error(err)
		}

		if res2 != test.want2 {
			t.Errorf("error")
		}

	}

}
