package utilities_test

import (
	"testing"
	"time"

	"github.com/derekpedersen/forget-me-please/utilities"
)

func TestDelay(t *testing.T) {
	// Arrange
	starttime := time.Now()

	// Act
	utilities.Delay()
	endtime := time.Now()

	// Assert
	if endtime.Sub(starttime) < time.Duration(time.Second*1) {
		t.Errorf("There was no delay!")
	}
	if endtime.Sub(starttime) > time.Duration(time.Second*11) {
		t.Errorf("The delay was too long!")
	}
}
