package rvalidation_test

import (
	"fmt"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rvalidation"
)

func TestValidateIPv4(t *testing.T) {
	type test struct {
		name             string
		hash             string
		expectedValidity bool
	}

	tests := []test{
		{name: "validIPv4_1", hash: "0.1.2.3", expectedValidity: true},
		{name: "validIPv4_2", hash: "172.16.1.2", expectedValidity: true},
		{name: "validIPv4_3", hash: "172.31.1.2", expectedValidity: true},
		{name: "validIPv4_4", hash: "192.167.1.2", expectedValidity: true},
		{name: "validIPv4_5", hash: "255.255.255.255", expectedValidity: true},
		{name: "InvalidIPv4_1", hash: "1.2.3.", expectedValidity: false},
		{name: "InvalidIPv4_2", hash: "1.2.256.4", expectedValidity: false},
		{name: "InvalidIPv4_3", hash: "1.256.3.4", expectedValidity: false},
		{name: "InvalidIPv4_4", hash: "1.2.3.4.5", expectedValidity: false},
		{name: "InvalidIPv4_5", hash: "1..3.4", expectedValidity: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValidHash := rvalidation.ValidateIPv4(tc.hash)
			if isValidHash != tc.expectedValidity {
				t.Errorf("ValidateIPv4(%v) Failed: expected %t, actual %t",
					tc.hash, tc.expectedValidity, isValidHash)
			}
		})
	}
}

func ExampleValidateIPv4() {
	fmt.Println(rvalidation.ValidateIPv4("192.168.1.2"))
	// Output: true
}
