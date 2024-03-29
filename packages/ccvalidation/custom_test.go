package ccvalidation

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestErrs_ValidateCCNumber(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "positive: American Express starts with 37",
			input:       "371449635398431",
			expectedErr: nil,
		},
		{
			name:        "positive: American Express starts with 34",
			input:       "340648704671028",
			expectedErr: nil,
		},
		{
			name:        "positive: Discover starts with 6011",
			input:       "6011634526224183",
			expectedErr: nil,
		},
		{
			name:        "positive: Discover starts with 644",
			input:       "6445644564456445",
			expectedErr: nil,
		},
		{
			name:        "positive: MasterCard starts with 51",
			input:       "5103747810889023",
			expectedErr: nil,
		},
		{
			name:        "positive: MasterCard starts with 2221",
			input:       "2221001234567896",
			expectedErr: nil,
		},
		{
			name:        "positive: Visa 16 length",
			input:       "4916664563529039",
			expectedErr: nil,
		},
		{
			name:        "negative: invalid type",
			input:       "1111111111111",
			expectedErr: errValidateCCNumberInvalidType,
		},
		{
			name:        "negative: American Express Luhn check failed",
			input:       "375533158339514",
			expectedErr: errValidateCCNumberLuhnCheckFailed,
		},
		{
			name:        "negative: Discover Luhn check failed",
			input:       "6011647480872146",
			expectedErr: errValidateCCNumberLuhnCheckFailed,
		},
		{
			name:        "negative: MasterCard Luhn check failed",
			input:       "5218858536048879",
			expectedErr: errValidateCCNumberLuhnCheckFailed,
		},
		{
			name:        "negative: Visa Luhn check failed",
			input:       "4916617520536879",
			expectedErr: errValidateCCNumberLuhnCheckFailed,
		},
		{
			name:        "negative: not numeric",
			input:       "test",
			expectedErr: errValidateCCNumberNumOnly,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ccValidator := NewCustomValidator()

			actual := ccValidator.ValidateCCNumber(tc.input)
			require.Equal(t, tc.expectedErr, actual, "errs are not equal")
		})
	}
}

func TestErrs_ValidateCCExpDate(t *testing.T) {
	month := fmt.Sprintf("%02d", time.Now().Month())
	year := fmt.Sprintf("%d", time.Now().Year()+1)

	t.Parallel()
	testCases := []struct {
		name        string
		input       ValidateCCExpDateOpt
		expectedErr error
	}{
		{
			name: "positive: month with leading zero",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  year,
			},
			expectedErr: nil,
		},
		{
			name: "positive: month without leading zero",
			input: ValidateCCExpDateOpt{
				Month: fmt.Sprintf("%d", time.Now().Month()),
				Year:  year,
			},
			expectedErr: nil,
		},
		{
			name: "positive: min value for month",
			input: ValidateCCExpDateOpt{
				Month: strconv.Itoa(monthMin),
				Year:  year,
			},
			expectedErr: nil,
		},
		{
			name: "positive: max value for month",
			input: ValidateCCExpDateOpt{
				Month: strconv.Itoa(monthMax),
				Year:  year,
			},
			expectedErr: nil,
		},
		{
			name: "positive: min value for year",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  strconv.Itoa(yearMin),
			},
			expectedErr: errValidateCCExpired,
		},
		{
			name: "positive: max value for year",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  strconv.Itoa(yearMax),
			},
			expectedErr: nil,
		},
		{
			name: "negative: empty month",
			input: ValidateCCExpDateOpt{
				Month: "",
				Year:  year,
			},
			expectedErr: errValidateCCExpDateMonthNumOnly,
		},
		{
			name: "negative: empty year",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  "",
			},
			expectedErr: errValidateCCExpDateYearNumOnly,
		},
		{
			name: "negative: month below of range",
			input: ValidateCCExpDateOpt{
				Month: fmt.Sprintf("%02d", monthMin-1),
				Year:  year,
			},
			expectedErr: errValidateCCExpDateMonthOutOfRange,
		},
		{
			name: "negative: month above of range",
			input: ValidateCCExpDateOpt{
				Month: fmt.Sprintf("%02d", monthMax+1),
				Year:  year,
			},
			expectedErr: errValidateCCExpDateMonthOutOfRange,
		},
		{
			name: "negative: year below of range",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  fmt.Sprintf("%d", yearMin-1),
			},
			expectedErr: errValidateCCExpDateYearOutOfRange,
		},
		{
			name: "negative: year above of range",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  fmt.Sprintf("%d", yearMax+1),
			},
			expectedErr: errValidateCCExpDateYearOutOfRange,
		},
		{
			name: "negative: not numeric month",
			input: ValidateCCExpDateOpt{
				Month: "test",
				Year:  year,
			},
			expectedErr: errValidateCCExpDateMonthNumOnly,
		},
		{
			name: "negative: not numeric year",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  "test",
			},
			expectedErr: errValidateCCExpDateYearNumOnly,
		},
		{
			name: "negative: expired",
			input: ValidateCCExpDateOpt{
				Month: month,
				Year:  fmt.Sprintf("%d", time.Now().Year()-1),
			},
			expectedErr: errValidateCCExpired,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ccValidator := NewCustomValidator()

			actual := ccValidator.ValidateCCExpDate(tc.input)
			require.Equal(t, tc.expectedErr, actual, "errs are not equal")
		})
	}
}
