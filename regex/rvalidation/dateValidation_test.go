package rvalidation_test

import (
	"fmt"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rvalidation"
)

func ExampleValidateDate() {
	fmt.Println(rvalidation.ValidateDate("21/01/1994"))
	// Output: true
}

func TestValidateDate(t *testing.T) {
	type test struct {
		name             string
		date             string
		expectedValidity bool
	}

	tests := []test{
		{name: "standardDate", date: "21/01/1994", expectedValidity: true},
		{name: "invalidDay", date: "31/11/2006", expectedValidity: false},
		{name: "leapYear", date: "29/02/2000", expectedValidity: true},
		{name: "notLeapYear", date: "29/02/2001", expectedValidity: false},
		{name: "invalidMonth", date: "29/14/2001", expectedValidity: false},
		{name: "invalidYear", date: "29/14/11", expectedValidity: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValidDate := rvalidation.ValidateDate(tc.date)
			if isValidDate != tc.expectedValidity {
				t.Errorf("ValidateDate(%v) Failed: expected %t, actual %t",
					tc.date, tc.expectedValidity, isValidDate)
			}
		})
	}
}
