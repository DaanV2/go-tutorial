package rand

import (
	"fmt"
	mrand "math/rand"

	"github.com/google/uuid"
)

// RandomID generates a random ID
func RandomID() string {
	n := mrand.Int63()
	// If over 2/3, generate a new UUID else generate a random hex
	if n > (1<<63)/3*2 {
		return fmt.Sprintf("%x", n)
	}

	return uuid.NewString()
}
