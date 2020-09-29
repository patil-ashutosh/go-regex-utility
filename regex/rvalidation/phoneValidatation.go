package rvalidation

import "regexp"

// 1) Mobile Number
// 			- Starts with 6,7,8,8
// 			- 10 digit
//			- Prexix can be +91, 0

// 2) LandLine Number
// 			- {area-code}-{local number}
// 			- 10 digit number
// 			- area codes range from 2-digits to 4-digits
// 			- ex. 02321-238200

//	pattern = `^(((?:(?:\+|0{0,2})91([\s-])?|[0]?)?[6-9]\d{9})|((0)?(([1-9]\d{1}\-\d{8})|([1-9]\d{2}-\d{7})|([1-9]\d{3}-\d{6}))))$`

// ValidatePhoneNumber validates mobile phone number and ladnline number in india
func ValidatePhoneNumber(num string) bool {

	prefixMobileNum := `(?:(?:\+|0{0,2})91([\s-])?|[0]?)?`
	mobileNum := `[6-9]\d{9}`
	landLineNum := `((0)?(([1-9]\d{1}-\d{8})|([1-9]\d{2}-\d{7})|([1-9]\d{3}-\d{6})))`
	pattern := "^(" + "(" + prefixMobileNum + mobileNum + ")" + "|" + landLineNum + ")$"

	//pattern = `^(((?:(?:\+|0{0,2})91([\s-])?|[0]?)?[6-9]\d{9})|((0)?(([1-9]\d{1}\-\d{8})|([1-9]\d{2}-\d{7})|([1-9]\d{3}-\d{6}))))$`

	reg := regexp.MustCompile(pattern)
	match := reg.MatchString(num)
	return match
}
