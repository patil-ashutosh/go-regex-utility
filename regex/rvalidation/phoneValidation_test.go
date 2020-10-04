package rvalidation_test

import (
	"fmt"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rvalidation"
)

func ExampleValidatePhoneNumber() {
	fmt.Println(rvalidation.ValidatePhoneNumber("+91-9819882936"))
	// Output: true
}

func TestValidatePhoneNumber(t *testing.T) {
	type test struct {
		name             string
		number           string
		expectedValidity bool
	}

	tests := []test{
		{name: "numberWithCountryCode", number: "+91-9819882936", expectedValidity: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValidNumber := rvalidation.ValidatePhoneNumber(tc.number)
			if isValidNumber != tc.expectedValidity {
				t.Errorf("ValidatePhoneNumber(%v) Failed: expected %t, actual %t",
					tc.number, tc.expectedValidity, isValidNumber)
			}
		})
	}
}
