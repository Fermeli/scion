package bloomfilter

import (
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"hash/crc64"
	"math"
	"time"

	"github.com/golang-collections/go-datastructures/bitarray"
)

// Data structure representing a bloom filter
type BloomFilter struct {
	nbOfHashes uint64
	bitArray   bitarray.BitArray
	length     uint64
}

// SlidingBloomFilter is a data structure representing multiple bloom filters
type SlidingBloomFilter struct {
	filters        []BloomFilter
	switchInterval time.Duration
	lastSwitch     time.Time
	currentIndex   uint
}

// NewBloomFilter creates a BloomFilter and returns it given the length
// of the bit array and the number of hashes that are perfomed
func NewBloomFilter(sizeOfBloomFilter uint64, nbOfHashes uint64) BloomFilter {
	return BloomFilter{nbOfHashes: nbOfHashes,
		bitArray: bitarray.NewBitArray(sizeOfBloomFilter), length: sizeOfBloomFilter}
}

// NewSlidingBloomFilter creates a SlidingBloomFilter struct. With the
// number of bloom filters and the time stamp validity period, the optimal
// switch interval is automatically calculated.
func NewSlidingBloomFilter(sizeOfBloomFilter uint64, nbOfHashes uint64, nbOfFilters uint,
	timeStampValidity time.Duration, now time.Time) SlidingBloomFilter {
	filters := make([]BloomFilter, 0)

	for i := 0; i < int(nbOfFilters); i += 1 {
		filters = append(filters, NewBloomFilter(sizeOfBloomFilter, nbOfHashes))
	}

	// For B bloom filters and D timestamp validity, the filters rotate every
	// D /(B - 1) nanoseconds so that added elements are stored for at least D nanoseconds
	switchInterval := timeStampValidity / (time.Duration(nbOfFilters) - 1)

	return SlidingBloomFilter{filters: filters, switchInterval: switchInterval,
		lastSwitch: now, currentIndex: 0}
}

// NewOptimalSlidingBloomFilter creates a SlidingBloomFilter with optimal bit array length
// and optimal number of hash functions. additionRate is in element/second,
// timeStampValidity is in nanoseconds.
func NewOptimalSlidingBloomFilter(falsePositiveRate float64, additionRate float64,
	timeStampValidity time.Duration, now time.Time, nbOfFilters uint) SlidingBloomFilter {

	falsePositiveRatePerFilter := 1 - math.Pow(1-falsePositiveRate, 1.0/float64(nbOfFilters))
	switchInterval := timeStampValidity / (time.Duration(nbOfFilters) - 1)
	maxNumberOfInsertedElementsPerFilter := math.Ceil(additionRate * math.Pow10(-9) * float64(switchInterval))
	lengthOfBitarray := math.Ceil(-maxNumberOfInsertedElementsPerFilter * math.Log(falsePositiveRatePerFilter) / math.Pow(math.Ln2, 2))
	numberOfHashes := math.Ceil(math.Ln2 * lengthOfBitarray / maxNumberOfInsertedElementsPerFilter)

	return NewSlidingBloomFilter(uint64(lengthOfBitarray), uint64(numberOfHashes), nbOfFilters,
		timeStampValidity, now)
}

// HasBeenSeen returns false if the identifier has not been seen over the last time stamp
// validity period. Otherwise, it returns true.
func (m *SlidingBloomFilter) HasBeenSeen(identifier []byte, now time.Time) (bool, error) {
	nbOfFilters := len(m.filters)

	if now.Sub(m.lastSwitch) >= m.switchInterval {
		m.currentIndex = (m.currentIndex + 1) % uint(nbOfFilters)
		m.filters[m.currentIndex].Reset()
		m.lastSwitch = now
	}

	hasBeenSeen := false
	indices, err := m.filters[m.currentIndex].ComputeHashes(identifier)

	if err != nil {
		return false, err
	}

	for i := nbOfFilters - 1; i >= 0; i -= 1 {
		//We check the current bloom filter at then end
		index := (m.currentIndex + uint(i)) % uint(nbOfFilters)
		filter := m.filters[index]
		// Only the bits of the current bloom filter are modified.
		seen, err := filter.CheckBitArray(indices, index == m.currentIndex)
		if err != nil {
			return false, err
		}

		if seen {
			hasBeenSeen = true
			break
		}
	}

	return hasBeenSeen, nil

}

// HasBeenSeen returns true if the function has already been called with this
// argument and false otherwise, with a false positive rate.
func (b *BloomFilter) HasBeenSeen(identifier []byte, setBits bool) (bool, error) {
	indices, err := b.ComputeHashes(identifier)

	if err != nil {
		return false, err
	}

	return b.CheckBitArray(indices, setBits)
}

// CheckBitArray checks if the bits of the array at the given indices are set to 1.
// If all these bits are set to one it returns true, otherwise it returns false.
// If the setBits parameter is true, then the bits at the given indices are set
// to one in the array, otherwise they are not modified.
func (b *BloomFilter) CheckBitArray(indices []uint64, setBits bool) (bool, error) {
	res := true
	for _, index := range indices {
		bitSet, err := b.bitArray.GetBit(index)
		if err != nil {
			return false, err
		}
		//if the bit has not been set we set it and we set the result to false
		if !bitSet {
			res = false

			if setBits {
				b.bitArray.SetBit(index)
			}

		}
	}
	return res, nil
}

// ComputeHashes computes the hashes of the identifier based on double hashes
// algorithm.
func (b *BloomFilter) ComputeHashes(identifier []byte) ([]uint64, error) {
	var res = make([]uint64, b.nbOfHashes)

	// We use two crc64 hash functions with a different polynomial to
	// produce the indices.
	hash1 := crc64.New(crc64.MakeTable(crc64.ISO))
	hash2 := crc64.New(crc64.MakeTable(crc64.ECMA))
	_, err := hash1.Write(identifier)
	if err != nil {
		return nil, err
	}
	_, err = hash2.Write(identifier)
	if err != nil {
		return nil, err
	}

	res1 := hash1.Sum64()
	res2 := hash2.Sum64()
	var i uint64

	for i = 0; i < b.nbOfHashes; i += 1 {
		res[i] = (res1 + i*res2) % b.length
	}

	return res, nil
}

// Reset sets all the bits to 0 in the bit array of the bloom filter
func (b *BloomFilter) Reset() {
	b.bitArray = bitarray.NewBitArray(b.length)
}
