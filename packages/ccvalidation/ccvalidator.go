package ccvalidation

type Validator interface {
	ValidateCCNumber(ccNumber string) error
	ValidateCCExpDate(opt ValidateCCExpDateOpt) error
}

type ValidateCCExpDateOpt struct {
	Month string
	Year  string
}
