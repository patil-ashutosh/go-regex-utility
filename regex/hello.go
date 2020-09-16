package regex

import (
	"fmt"
)

// Hello sample program
func Hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude"
	}
	return fmt.Sprintf("Hello %v!", user)
}
