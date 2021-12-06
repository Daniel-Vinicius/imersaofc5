package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, ExpirationCVV int) (*CreditCard, error) {
	creditCardCreated := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             ExpirationCVV,
	}

	err := creditCardCreated.IsValid()

	if err != nil {
		return nil, err
	}

	return creditCardCreated, nil
}

func (creditCard *CreditCard) IsValid() error {
	err := creditCard.validateNumber()
	if err != nil {
		return err
	}

	err = creditCard.validateMonth()
	if err != nil {
		return err
	}

	err = creditCard.validateYear()
	if err != nil {
		return err
	}

	return nil
}

func (creditCard *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !re.MatchString(creditCard.number) {
		return errors.New("invalid credit card number")
	}

	return nil
}

func (creditCard *CreditCard) validateMonth() error {
	if creditCard.expirationMonth > 0 && creditCard.expirationMonth < 13 {
		return nil
	}

	return errors.New("invalid expiration month")
}

func (creditCard *CreditCard) validateYear() error {
	currentYear := time.Now().Year()

	if creditCard.expirationYear >= currentYear {
		return nil
	}

	return errors.New("invalid expiration year")
}
