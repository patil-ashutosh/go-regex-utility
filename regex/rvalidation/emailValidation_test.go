package rvalidation_test

import (
	"fmt"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rvalidation"
)

func TestValidateEmail(t *testing.T) {
	type test struct {
		name             string
		email            string
		expectedValidity bool
	}

	tests := []test{
		{name: "validEmail1", email: "john@gmail.com", expectedValidity: true},
		{name: "validEmail2", email: "johaan@gmail.com", expectedValidity: true},
		{name: "InvalidEmailWithSpecialCharAtStart", email: "_john@gmail.com", expectedValidity: false},
		{name: "InvalidEmailWithSpecialCharAtEnd", email: "john_@gmail.com", expectedValidity: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValidEmail := rvalidation.ValidateEmail(tc.email)
			if isValidEmail != tc.expectedValidity {
				t.Errorf("ValidateEmail(%v) Failed: expected %t, actual %t",
					tc.email, tc.expectedValidity, isValidEmail)
			}
		})
	}
}

func ExampleValidateEmail() {
	fmt.Println(rvalidation.ValidateEmail("sanajy@gmail.com"))
	// Output: true
}
