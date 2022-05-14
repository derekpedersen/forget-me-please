package twitter_test

import (
	"testing"

	"github.com/derekpedersen/forget-me-please/twitter"
)

func TestNewUser(t *testing.T) {
	// Arrange
	config := twitter.Config{
		UserName: "Derek",
	}

	// Act
	_, err := twitter.NewUser(config)

	// Assert
	if err != nil {
		t.Fatalf("Encountered unexpected error: %v", err)
	}
	// if user.Data.UserName != "Derek" {
	// 	t.Errorf("Expected UserName to be %v but was %v", "Derek", user.Data.UserName)
	// }
}
