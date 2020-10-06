package rvalidation

import (
	"fmt"
	"regexp"
)

// ValidateIPv4 validates a string as an ipv4 address.
func ValidateIPv4(hash string) bool {
	class := `([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`
	pattern := fmt.Sprintf(`^%[1]s\.%[1]s\.%[1]s\.%[1]s$`, class)
	reg := regexp.MustCompile(pattern)
	match := reg.MatchString(hash)

	return match
}
