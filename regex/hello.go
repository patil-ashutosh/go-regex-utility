package regex

import (
	"fmt"
)

func Hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude"
	}
	return fmt.Sprintf("Hello %v!", user)
}
