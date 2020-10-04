package rvalidation_test

import (
	"fmt"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rvalidation"
)

func TestValidateMd5Hash(t *testing.T) {
	type test struct {
		name             string
		hash             string
		expectedValidity bool
	}

	tests := []test{
		{name: "validMd5Hash1", hash: "9a41402abc4b2a76b9719d911017c592", expectedValidity: true},
		{name: "validMd5Hash2", hash: "9f6e6800cfae7749eb6c486619254b9c", expectedValidity: true},
		{name: "validMd5Hash3", hash: "7eae4e3fe6ef4391a202e274842830a4", expectedValidity: true},
		{name: "validMd5Hash4", hash: "028743a2ea021ef2d48ad3d21d29ea1a", expectedValidity: true},
		{name: "validMd5Hash5", hash: "952bccf9afe8e4c04306f70f7bed6610", expectedValidity: true},
		{name: "InvalidMd5HashLength", hash: "952bcczzz9afe8e4c04306f70f7bed6610", expectedValidity: false},
		{name: "validMd5Hash5", hash: "952bccf9afe8e4c04306f70f7bed6610", expectedValidity: true},
		{name: "validMd5Hash5", hash: "952bccf9afe8e4c04306f70f7bed6610", expectedValidity: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValidHash := rvalidation.ValidateMd5Hash(tc.hash)
			if isValidHash != tc.expectedValidity {
				t.Errorf("ValidateMd5Hash(%v) Failed: expected %t, actual %t",
					tc.hash, tc.expectedValidity, isValidHash)
			}
		})
	}
}

func ExampleValidateMd5Hash() {
	fmt.Println(rvalidation.ValidateMd5Hash("952bccf9afe8e4c04306f70f7bed6610"))
	// Output: true
}
