package ccvalidation

import (
	"darkness8129/validation-service/packages/errs"
	"regexp"
)

// Validator provides a logic for credit card validation.
// A credit card is considered invalid if any method returns an error; otherwise, it is considered valid.
type Validator interface {
	ValidateCCNumber(ccNumber string) error
	ValidateCCExpDate(opt ValidateCCExpDateOpt) error
}

type ValidateCCExpDateOpt struct {
	Month string
	Year  string
}

const (
	numOnlyErrCode         = "num_only"
	invalidTypeErrCode     = "invalid_type"
	luhnCheckFailedErrCode = "luhn_check_failed"
	expiredErrCode         = "expired"
	outOfRange             = "out_of_range"
)

var (
	errValidateCCNumberNumOnly         = errs.New(errs.Options{Message: "cc number must contain only numbers", Code: numOnlyErrCode})
	errValidateCCNumberInvalidType     = errs.New(errs.Options{Message: "invalid cc type", Code: invalidTypeErrCode})
	errValidateCCNumberLuhnCheckFailed = errs.New(errs.Options{Message: "Luhn check failed", Code: luhnCheckFailedErrCode})

	errValidateCCExpDateMonthNumOnly    = errs.New(errs.Options{Message: "exp month must contain only numbers", Code: numOnlyErrCode})
	errValidateCCExpDateYearNumOnly     = errs.New(errs.Options{Message: "exp year must contain only numbers", Code: numOnlyErrCode})
	errValidateCCExpDateMonthOutOfRange = errs.New(errs.Options{Message: "exp month out of range", Code: outOfRange})
	errValidateCCExpDateYearOutOfRange  = errs.New(errs.Options{Message: "exp year out of range", Code: outOfRange})
	errValidateCCExpired                = errs.New(errs.Options{Message: "cc is expired", Code: expiredErrCode})
)

type ccType string

// Supported types of credit cards by validator.
var (
	ccTypeAmericanExpress ccType = "American Express"
	ccTypeDiscover        ccType = "Discover"
	ccTypeMasterCard      ccType = "MasterCard"
	ccTypeVisa            ccType = "Visa"
)

var ccTypeRegexps = []struct {
	Type   ccType
	Regexp *regexp.Regexp
}{
	{
		Type:   ccTypeAmericanExpress,
		Regexp: regexp.MustCompile(`^3[47][0-9]{13}$`),
	},
	{
		Type:   ccTypeDiscover,
		Regexp: regexp.MustCompile(`^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$`),
	},
	{
		Type:   ccTypeMasterCard,
		Regexp: regexp.MustCompile(`^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$`),
	},
	{
		Type:   ccTypeVisa,
		Regexp: regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
	},
}
