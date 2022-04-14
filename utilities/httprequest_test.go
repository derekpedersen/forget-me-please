package utilities_test

import (
	"net/http"
	"testing"

	"github.com/derekpedersen/forget-me-please/utilities"
)

func TestHttpRequest(t *testing.T) {
	// Arrange
	url := "https://google.com"

	// Act
	response, err := utilities.HttpRequest(url, http.MethodGet, nil)

	// Assert
	if err != nil {
		t.Fatalf("An unexpected error: %v", err)
	}
	if response == nil {
		t.Fatalf("No response was retured!")
	}
	if len(*response) == 0 {
		t.Errorf("Expected it to contain a response")
	}
}
