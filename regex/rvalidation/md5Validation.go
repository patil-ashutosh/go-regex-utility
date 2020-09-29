package rvalidation

import "regexp"

// ValidateMd5Hash validates a string as an md5 hash.
func ValidateMd5Hash(hash string) bool {
	pattern := `^([a-f0-9]{32})$`
	reg := regexp.MustCompile(pattern)
	match := reg.MatchString(hash)
	return match

}
