package ccvalidation

import "darkness8129/validation-service/packages/errs"

type Validator interface {
	ValidateCCNumber(ccNumber string) error
	ValidateCCExpDate(opt ValidateCCExpDateOpt) error
}

const (
	numOnlyErrCode = "num_only"
	expiredErrCode = "expired"
	outOfRange     = "out_of_range"
)

var (
	errValidateCCNumberNumOnly = errs.New(errs.Options{Message: "cc number must contain only numbers", Code: numOnlyErrCode})

	errValidateCCExpDateMonthNumOnly    = errs.New(errs.Options{Message: "exp month must contain only numbers", Code: numOnlyErrCode})
	errValidateCCExpDateYearNumOnly     = errs.New(errs.Options{Message: "exp year must contain only numbers", Code: numOnlyErrCode})
	errValidateCCExpDateMonthOutOfRange = errs.New(errs.Options{Message: "exp month out of range", Code: outOfRange})
	errValidateCCExpDateYearOutOfRange  = errs.New(errs.Options{Message: "exp year out of range", Code: outOfRange})
	errValidateCCExpired                = errs.New(errs.Options{Message: "cc is expired", Code: expiredErrCode})
)

type ValidateCCExpDateOpt struct {
	Month string
	Year  string
}
