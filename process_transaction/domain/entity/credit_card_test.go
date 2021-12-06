package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("0000000000000000", "Jose da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5457668645808073", "Jose da Silva", 12, 2024, 509)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5457668645808073", "Jose da Silva", 13, 2024, 509)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5457668645808073", "Jose da Silva", 0, 2024, 509)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5457668645808073", "Jose da Silva", 11, 2024, 509)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	nextYear := time.Now().AddDate(1, 0, 0)

	_, err := NewCreditCard("5457668645808073", "Jose da Silva", 12, lastYear.Year(), 509)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5457668645808073", "Jose da Silva", 11, nextYear.Year(), 509)
	assert.Nil(t, err)
}
