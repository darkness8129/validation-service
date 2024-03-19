package service

import (
	"context"
	"darkness8129/validation-service/packages/ccvalidation"
	"darkness8129/validation-service/packages/errs"
	"darkness8129/validation-service/packages/logging"
	"fmt"
)

var _ CreditCardService = (*creditCardService)(nil)

type creditCardService struct {
	logger    logging.Logger
	validator ccvalidation.Validator
}

func NewCreditCardService(logger logging.Logger, validator ccvalidation.Validator) *creditCardService {
	return &creditCardService{
		logger:    logger.Named("creditCardService"),
		validator: validator,
	}
}

func (s *creditCardService) Validate(ctx context.Context, opt ValidateOpt) error {
	logger := s.logger.Named("Validate")

	err := s.validator.ValidateCCNumber(opt.CCNumber)
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return err
		}

		logger.Error("failed to validate card number", "err", err)
		return fmt.Errorf("failed to validate card number: %w", err)
	}
	logger.Debug("validated card number")

	err = s.validator.ValidateCCExpDate(ccvalidation.ValidateCCExpDateOpt{
		Month: opt.CCExpMonth,
		Year:  opt.CCExpYear,
	})
	if err != nil {
		if errs.IsCustom(err) {
			logger.Info(err.Error())
			return err
		}

		logger.Error("failed to validate expiration date", "err", err)
		return fmt.Errorf("failed to validate expiration date: %w", err)
	}

	logger.Info("credit card successfully validated")
	return nil
}
