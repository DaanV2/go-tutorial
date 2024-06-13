package rand_test

import (
	"testing"
	"time"

	"github.com/DaanV2/go-tutorial/pkg/rand"
)

func TestRandomTime(t *testing.T) {
	expectedFormat := "2006-01-02T15:04:05Z07:00" // Replace with the expected format

	// Call the function
	result := rand.RandomTime()

	// Parse the result into a time.Time value
	parsedTime, err := time.Parse(time.RFC3339, result)
	if err != nil {
		t.Errorf("Failed to parse the random time: %v", err)
	}

	// Check if the parsed time matches the expected format
	if parsedTime.Format(expectedFormat) != result {
		t.Errorf("Unexpected time format. Expected: %s, Got: %s", expectedFormat, result)
	}
}