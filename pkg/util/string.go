package util

import "math/rand"

var randx = rand.NewSource(42)

// RandString returns a random hex string of length n.
func RandString(n int) string {
	const hexBytes = "0123456789ABCDEF"
	const (
		hexIdxBits = 4                 // 4 bits to represent a hex string
		hexIdxMask = 1<<hexIdxBits - 1 // All 1-bits, as many as hexIdxBits
		hexIdxMax  = 63 / hexIdxBits   // # of hex indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for hexIdxMax letters!
	for i, cache, remain := n-1, randx.Int63(), hexIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randx.Int63(), hexIdxMax
		}
		if idx := int(cache & hexIdxMask); idx < len(hexBytes) {
			b[i] = hexBytes[idx]
			i--
		}
		cache >>= hexIdxBits
		remain--
	}

	return string(b)
}
