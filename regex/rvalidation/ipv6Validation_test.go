package rvalidation_test

import (
	"fmt"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rvalidation"
)

func TestValidateIpv6(t *testing.T) {
	type test struct {
		name             string
		address          string
		expectedValidity bool
	}

	tests := []test{
		{name: "standardFormat", address: "ABCD:EF01:2345:6789:ABCD:EF01:2345:6789", expectedValidity: true},
		{name: "standardFormatMixed", address: "ABCD:EF01:2345:6789:ABCD:EF01:2345:6789:172.16.1.2", expectedValidity: true},
		{name: "standardFormatNoLeadingZeros", address: "2001:DB8:0:0:8:800:200C:417A", expectedValidity: true},
		{name: "standardFormatUnspecified", address: "0:0:0:0:0:0:0:0", expectedValidity: true},
		{name: "standardFormatMixedCase", address: "AbCd:Ef01:2345:9:AeaD:E01:2345:6789", expectedValidity: true},
		{name: "compressedFormat0", address: "2001:DB8:ABCD:8:800:200C:A::", expectedValidity: true},
		{name: "compressedFormat0WithIpv4", address: "2001:DB8:ABCD:8:800:200C:A::1.2.3.4", expectedValidity: true},
		{name: "compressedFormat1NoEnd", address: "2001:DB8:ABCD:8:800:200C::", expectedValidity: true},
		{name: "compressedFormat1End", address: "2001:DB8:ABCD:8:800:200C::12bc", expectedValidity: true},
		{name: "compressedFormat1EndWithIpv4", address: "2001:DB8:ABCD:8:800:200C::12bc:1.2.3.4", expectedValidity: true},
		{name: "compressedFormat2NoEnd", address: "2001:DB8:ABCD:8:800::", expectedValidity: true},
		{name: "compressedFormat2NoEndWithIpv4", address: "2001:DB8:ABCD:8:800::1.2.3.4", expectedValidity: true},
		{name: "compressedFormat2End", address: "2001:DB8:ABCD:8:800::12Ac", expectedValidity: true},
		{name: "compressedFormat3NoEnd", address: "2001:DB8:ABCD:8::", expectedValidity: true},
		{name: "compressedFormat3End", address: "2001:DB8:ABCD:800::12Ac", expectedValidity: true},
		{name: "compressedFormat3EndWithIpv4", address: "2001:DB8:ABCD:800::12Ac:172.0.0.1", expectedValidity: true},
		{name: "compressedFormat4NoEnd", address: "2001:DB8:8::", expectedValidity: true},
		{name: "compressedFormat4NoEndWithIpv4", address: "2001:DB8:8::172.168.1.1", expectedValidity: true},
		{name: "compressedFormat4End", address: "2001:DB8:800::12Ac:12af", expectedValidity: true},
		{name: "compressedFormat5NoEnd", address: "2001:DB8::", expectedValidity: true},
		{name: "compressedFormat5End", address: "2001:800::12Ac:12af", expectedValidity: true},
		{name: "compressedFormat5NoEndWithIpv4", address: "2001:DB8::192.168.1.1", expectedValidity: true},
		{name: "compressedFormat6NoEnd", address: "DB8::", expectedValidity: true},
		{name: "compressedFormat6End", address: "800::3984:12Ac:12af", expectedValidity: true},
		{name: "compressedFormat6EndWithIpv4", address: "800::3984:12Ac:12af:192.1.1.0", expectedValidity: true},
		{name: "compressedFormat7NoEnd", address: "::", expectedValidity: true},
		{name: "compressedFormat7NoEndWithIpv4", address: "::13.1.68.3", expectedValidity: true},
		{name: "compressedFormat7End", address: "::12af:3244:2312:a34b:A122:a123:1244", expectedValidity: true},
		{name: "invalidCompressedFormat7End", address: "::12af:3244::2312:a34b:A122:a123:1244", expectedValidity: false},
		{name: "invalidCompressedFormat7EndWithIpv4", address: "::1244.3.45.55a", expectedValidity: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValidIpv6 := rvalidation.ValidateIpv6(tc.address)
			if isValidIpv6 != tc.expectedValidity {
				t.Errorf("ValidateIpv6(%v) Failed: expected %t, actual %t",
					tc.address, tc.expectedValidity, isValidIpv6)
			}
		})
	}
}

func ExampleValidateIpv6() {
	fmt.Println(rvalidation.ValidateIpv6("::1"))
	// Output: true
}
