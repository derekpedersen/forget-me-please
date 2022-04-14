package utilities_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/derekpedersen/forget-me-please/utilities"
)

func TestReadline(t *testing.T) {
	// Arrange
	input := "Derek\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// Act
	line := utilities.ReadLine(reader)

	// Assert
	if line == nil {
		t.Fatalf("No input was read!")
	}
	if *line != "Derek" {
		t.Errorf("Unable to read line correctly! Expected %v got %v", "Derek", line)
	}
}
