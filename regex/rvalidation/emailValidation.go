package rvalidation

import "regexp"

// https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address

// A valid email address has four parts:
//     1) Username
// 				- Uppercase and lowercase letters
// 				- Digits from 0 to 9
// 				- Special characters such as ! # $ % & ' * + - / = ? ^ _ ` { |
// 				- A special character cannot appear as the first or last character
//     2) @ symbol
//     3) Domain name
// 				- Domain names may be a maximum of 255 characters and consist of:
// 				- Uppercase and lowercase letters in English (A-Z, a-z)
// 				- Digits from 0 to 9
// 				- A hyphen (-)
//				- A period (.)  (used to identify a sub-domain; for example,  email.domainsample)
//     4) Top-level domain

// Email -> userName@domainName.topLevelDomainName
// [a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}    Username,
// [a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])? Domain , First and last char can be only letter and number
// (?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*   Top-level domain
// pattern = "^[a-zA-Z0-9][a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{0,62}[a-zA-Z0-9]@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,253}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func ValidateEmail(email string) bool {

	userName := "[a-zA-Z0-9][a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{0,62}[a-zA-Z0-9]"
	domainName := "[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,253}[a-zA-Z0-9])?"
	topLevelDomainName := "(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*"
	pattern := "^" + userName + "@" + domainName + topLevelDomainName + "$"

	reg := regexp.MustCompile(pattern)
	match := reg.MatchString(email)

	return match
}
