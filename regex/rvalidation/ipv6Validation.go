// Package rvalidation is collection of different validation scripts using regex
package rvalidation

import (
	"fmt"
	"regexp"
)

// ValidateIpv6 method validates Ipv6 network address and returns true or false
//
// ________________________________________________________________________________________
//
// A Valid ipv6 address as defined by RFC 4291 (https://tools.ietf.org/html/rfc4291) can be in the following formats
//
//     1) The standard form is x:x:x:x:x:x:x:x, where the 'x's are one to
//        four hexadecimal digits of the eight 16-bit pieces of the address.
//        Examples:
//
//          ABCD:EF01:2345:6789:ABCD:EF01:2345:6789
//
//          2001:DB8:0:0:8:800:200C:417A
//
//		  Note that the leading zeros are optional, but every 16-bit piece should contain atleast one digit
//		  with the exception of scenario mentioned in 2.
//
//     2) An ipv6 address can be in a compressed form.
//		  This format uses "::" to indicate one or more groups of 16 bits of zeros.
//        The "::" can only appear once in an address.
//		  The "::" can also be used to compress leading or trailing zeros
//
//			For example, the following addresses
//			2001:DB8:0:0:8:800:200C:417A   a unicast address
//			FF01:0:0:0:0:0:0:101           a multicast address
//			0:0:0:0:0:0:0:1                the loopback address
//			0:0:0:0:0:0:0:0                the unspecified address
//
//			may be represented as
//
//			2001:DB8::8:800:200C:417A      a unicast address
//			FF01::101                      a multicast address
//			::1                            the loopback address
//			::                             the unspecified address
//
//     3) Additionally, in mixed format addressing there can be a combination of ipv4 and ipv6 addressing.
//		  The format is x:x:x:x:x:x:d.d.d.d, where the 'x's are the hexadecimal values of the 16-bit pieces
//        of the address, and the 'd's are the decimal values of the four low-order 8-bit pieces of the
//	      address
//
//			Examples:
//			0:0:0:0:0:0:13.1.68.3
//			0:0:0:0:0:FFFF:129.144.52.38
//
//			or in compressed form:
//
//			::13.1.68.3
//			::FFFF:129.144.52.38
//
// ________________________________________________________________________________________
//
func ValidateIpv6(address string) bool {
	// We check against: Standard format, standard+mixed format , each variation of the compressed format,
	// and each variation of compressed format + mixed format
	// Standard Format
	// a standard format ipv6 will have 7 of the 16-bit pieces
	standardPiece := "(?:[a-fA-F0-9]{1,4}:)"
	// a standard format ipv6 will have one end of the 16bit piece
	standardEndAnchor := "[a-fA-F0-9]{1,4}$"

	bit8Piece := `([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`
	ipv4EndAnchor := fmt.Sprintf(`%[1]s\.%[1]s\.%[1]s\.%[1]s$`, bit8Piece)

	standardPattern := "^" + standardPiece + "{7}" + standardEndAnchor
	standardPatternWithIpv4 := "^" + standardPiece + "{8}" + ipv4EndAnchor

	// Compressed Format
	// a compressed format ipv6 will have at most one "::" to indicate one or more groups of 16 bits of zeros.
	// if an address has a "::" notation there can be at-most 7 other ":" before and after this notation

	// when address has 7 pieces to left of ::
	compressedPattern0 := "^" + standardPiece + "{7}" + ":$"
	compressedPattern0WithIpv4 := "^" + standardPiece + "{7}" + ":" + ipv4EndAnchor

	// when address has 6 pieces to left of ::
	compressedPattern1 := "^" + standardPiece + "{6}" + ":" + "((" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern1WithIpv4 := "^" + standardPiece + "{6}" + ":" + "((" + standardPiece + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	// when address has 5 pieces to left of ::
	compressedPattern2 := "^" + standardPiece + "{5}" + ":" + "((" + standardPiece + "{0,1}" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern2WithIpv4 := "^" + standardPiece + "{5}" + ":" + "((" + standardPiece + "{0,2}" + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	// when address has 4 pieces to left of ::
	compressedPattern3 := "^" + standardPiece + "{4}" + ":" + "((" + standardPiece + "{0,2}" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern3WithIpv4 := "^" + standardPiece + "{4}" + ":" + "((" + standardPiece + "{0,3}" + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	// when address has 3 pieces to left of ::
	compressedPattern4 := "^" + standardPiece + "{3}" + ":" + "((" + standardPiece + "{0,3}" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern4WithIpv4 := "^" + standardPiece + "{3}" + ":" + "((" + standardPiece + "{0,4}" + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	// when address has 2 pieces to left of ::
	compressedPattern5 := "^" + standardPiece + "{2}" + ":" + "((" + standardPiece + "{0,4}" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern5WithIpv4 := "^" + standardPiece + "{2}" + ":" + "((" + standardPiece + "{0,5}" + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	// when address has 1 piece to left of ::
	compressedPattern6 := "^" + standardPiece + ":" + "((" + standardPiece + "{0,5}" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern6WithIpv4 := "^" + standardPiece + ":" + "((" + standardPiece + "{0,6}" + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	// when address has no piece to left of ::
	compressedPattern7 := "^" + "::" + "((" + standardPiece + "{0,6}" + standardEndAnchor + ")" + "|" + "$)"
	compressedPattern7WithIpv4 := "^" + "::" + "((" + standardPiece + "{0,7}" + ipv4EndAnchor + ")" + "|" + ipv4EndAnchor + ")"

	patterns := []string{}
	patterns = append(patterns, standardPattern, standardPatternWithIpv4)
	patterns = append(patterns, compressedPattern0, compressedPattern0WithIpv4)
	patterns = append(patterns, compressedPattern1, compressedPattern1WithIpv4)
	patterns = append(patterns, compressedPattern2, compressedPattern2WithIpv4)
	patterns = append(patterns, compressedPattern3, compressedPattern3WithIpv4)
	patterns = append(patterns, compressedPattern4, compressedPattern4WithIpv4)
	patterns = append(patterns, compressedPattern5, compressedPattern5WithIpv4)
	patterns = append(patterns, compressedPattern6, compressedPattern6WithIpv4)
	patterns = append(patterns, compressedPattern7, compressedPattern7WithIpv4)

	for _, pattern := range patterns {
		reg := regexp.MustCompile(pattern)
		if match := reg.MatchString(address); match {
			return match
		}
	}

	return false
}
