package ccvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var _ Validator = (*customValidator)(nil)

type customValidator struct{}

func NewCustomValidator() *customValidator {
	return &customValidator{}
}

// Limits of valid values for month and year
const (
	monthMax = 12
	monthMin = 1
	yearMax  = 9999
	yearMin  = 1000
)

// 1. The cc number is numeric.
// 2. The cc number matches with one of valid types.
// 3. The cc number passed Luhn Algo check.
func (s *customValidator) ValidateCCNumber(ccNumber string) error {
	var numericRegex = regexp.MustCompile(`^\d+$`)
	if !numericRegex.MatchString(ccNumber) {
		return errValidateCCNumberNumOnly
	}

	validType := false
	for _, v := range ccTypeRegexps {
		if v.Regexp.MatchString(ccNumber) {
			validType = true
			break
		}
	}
	if !validType {
		return errValidateCCNumberInvalidType
	}

	luhnCheckPassed, err := luhnCheck(ccNumber)
	if err != nil {
		return fmt.Errorf("failed to make Luhn check: %w", err)
	}
	if !luhnCheckPassed {
		return errValidateCCNumberLuhnCheckFailed
	}

	return nil
}

// 1. Loop over the digits of the card number backwards.
// 2. Multiply every second digit by two.
// 3. For any result that's 10 or more, subtract 9 to find the sum.
// 4. The card is valid if the total of all processed digits is divisible by 10.
func luhnCheck(ccNumber string) (bool, error) {
	sum := 0
	isSecondDigit := false

	for i := len(ccNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(ccNumber[i]))
		if err != nil {
			return false, fmt.Errorf("failed to convert str digit to int: %w", err)
		}

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isSecondDigit = !isSecondDigit
	}

	return sum%10 == 0, nil
}

// 1. The month is numeric and within the valid range.
// 2. The year is numeric and within the valid range.
// 3. The card is not expired.
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
