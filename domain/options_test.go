package domain_test

import (
	"testing"

	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
)

func TestPrintOptions(t *testing.T) {
	// Arrange
	opts := domain.Options(make(map[string]model.Option))

	// Act
	opts.PrintOptions()

	// Assert
	// TODO: how do I assert they owe test?
}

func TestSelectOption(t *testing.T) {
	// Arrange
	opts := domain.Options(make(map[string]model.Option))
	opts["Derek"] = model.Option{
		Key: "Derek",
	}

	// Act
	opt := opts.SelectOption("Derek")

	// Assert
	if opt != nil {
		t.Fatalf("No option was returned!")
	}
}
