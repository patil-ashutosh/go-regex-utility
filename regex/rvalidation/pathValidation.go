package rvalidation

import "regexp"

// If you want to add some another type of OS, you can just add
// the related path regex to this array
var regexes = []string{"^(\\/[a-zA-Z0-9_-]+)+|\\/$"}

// ValidatePath method validates paths and return true of false
//
// This function supports Linux-based and Windows operating systems
// ________________________________________________________________________________________
//
//    Linux-based Examples:
//			- /
//			- /dir1
//			- /dir1/
//			- /dir1/dir2
//
//	  Windows Examples:
//			- C://
//

func ValidatePath(path string) bool {
	for _, regex := range regexes {
		reg := regexp.MustCompile(regex)
		match := reg.MatchString(path)
		if match {
			return true
		}
	}
	return false
}
