package rand

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomHex generates a random hex string of length n
func RandomHex(n int) string {
	// Prepare a byte slice of half the length (since each byte is 2 hex characters)
	d := make([]byte, n/2)
	// Read random bytes into the slice
	_, _ = rand.Read(d)

	// Return the hex representation of the byte slice
	return hex.EncodeToString(d)
}
