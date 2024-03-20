package ccvalidation

import (
	"regexp"
	"strconv"
	"time"
)

var _ Validator = (*customValidator)(nil)

type customValidator struct{}

func NewCustomValidator() *customValidator {
	return &customValidator{}
}

const (
	monthMax = 12
	monthMin = 1
	yearMax  = 9999
	yearMin  = 1000
)

func (s *customValidator) ValidateCCNumber(ccNumber string) error {
	var numericRegex = regexp.MustCompile(`^\d+$`)
	if !numericRegex.MatchString(ccNumber) {
		return errValidateCCNumberNumOnly
	}

	return nil
}

func (s *customValidator) ValidateCCExpDate(opt ValidateCCExpDateOpt) error {
	month, err := strconv.Atoi(opt.Month)
	if err != nil {
		return errValidateCCExpDateMonthNumOnly
	}
	if month > monthMax || month < monthMin {
		return errValidateCCExpDateMonthOutOfRange
	}

	year, err := strconv.Atoi(opt.Year)
	if err != nil {
		return errValidateCCExpDateYearNumOnly
	}
	if year > yearMax || year < yearMin {
		return errValidateCCExpDateYearOutOfRange
	}

	// get beginning (00:00) of the first day of the next month,
	// because the credit card expires on the first day of the month
	// following the one indicated on the card
	nextMonth := time.Month(month) + 1
	expDate := time.Date(year, nextMonth, 1, 0, 0, 0, 0, time.UTC)
	if time.Now().UTC().After(expDate) {
		return errValidateCCExpired
	}

	return nil
}
