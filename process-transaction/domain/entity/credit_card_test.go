package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	creditCard, err := NewCreditCard("4000000000000000", "José da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())
}
