package ccvalidation

var _ Validator = (*customValidator)(nil)

type customValidator struct{}

func NewCustomValidator() *customValidator {
	return &customValidator{}
}

func (s *customValidator) ValidateCCNumber(ccNumber string) error {
	return nil
}

func (s *customValidator) ValidateCCExpDate(opt ValidateCCExpDateOpt) error {
	return nil
}
