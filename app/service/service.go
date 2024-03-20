package service

import "context"

type Services struct {
	CreditCard CreditCardService
}

// HTTPServer provides a logic for interacting with credit card.
type CreditCardService interface {
	Validate(ctx context.Context, opt ValidateOpt) error
}

type ValidateOpt struct {
	CCNumber   string
	CCExpMonth string
	CCExpYear  string
}
