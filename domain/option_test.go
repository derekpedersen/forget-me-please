package domain_test

import (
	"testing"

	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
)

func TestProcessOption(t *testing.T) {
	// Arrange
	opt := domain.Option(model.Option{})

	// Act
	domain.ProcessOption(&opt)

	// Assert
	// TODO: how do I assert they owe test?
}
