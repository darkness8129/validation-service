package service

import "context"

type Services struct {
	CreditCard CreditCardService
}

type CreditCardService interface {
	Validate(ctx context.Context, opt ValidateOpt) error
}

type ValidateOpt struct {
	CCNumber   string
	CCExpMonth string
	CCExpYear  string
}
