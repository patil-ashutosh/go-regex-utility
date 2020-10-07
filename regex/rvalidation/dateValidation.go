package rvalidation

import (
	"regexp"
)

// ValidateDate validates date strings of format dd/mm/yyyy.
func ValidateDate(hash string) bool {
	regularDay := "(0[1-9]|1[0-9]|2[0-8])"
	regularMonth := "(0[1-9]|1[012])"
	regularMonthEnds := `((29|30|31)[\/](0[13578]|1[02]))|((29|30)[\/](0[4,6,9]|11))`
	regularDayMonth := `((` + regularDay + `[\/]` + regularMonth + `)|` + regularMonthEnds + `)`
	regularYear := `(19|[2-9][0-9])\d\d`

	regularDatePattern := `(^` + regularDayMonth + `[\/]` + regularYear + `$)`

	leapDay := `^29[\/]02`
	leapYears := "(19|[2-9][0-9])(00|04|08|12|16|20|24|28|32|36|40|44|48|52|56|60|64|68|72|76|80|84|88|92|96)"
	leapDatePattern := `(` + leapDay + `[\/]` + leapYears + `$)`

	pattern := regularDatePattern + "|" + leapDatePattern

	reg := regexp.MustCompile(pattern)

	return reg.MatchString(hash)
}
