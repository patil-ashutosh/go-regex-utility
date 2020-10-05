// Package rstring is collection of various string programs using regex
package rstring

import (
	"regexp"
)

// CountStringOccurrenceInString Count the number of times a particular string occurs in another string.
func CountStringOccurrenceInString(input string, subString string) int {
	re := regexp.MustCompile(subString)
	results := re.FindAllString(input, -1)

	return len(results)
}

// SplitString splits string into a slices based on separator(delimiter, line break, regex, etc.).
//
//  If Separator is not provided then Split String on Whitespace
//
//  strs take 2 arguments,  separator string and string  which need to be split.
//  int n , is max split count (use n = -1, for max split)
//  n > 0: at most n substrings; the last substring will be the unsplit remainder.
//  n == 0: the result is nil (zero substrings)
//  n < 0: all substrings
func SplitString(n int, strs ...string) []string {
	str := ""
	sep := ""

	const inputStringLen = 2

	if len(strs) == 1 {
		str = strs[0]
		sep = "\\s+"
	} else if len(strs) == inputStringLen {
		str = strs[0]
		sep = strs[1]
	}

	re := regexp.MustCompile(sep)
	split := re.Split(str, n)

	return split
}

// ContainsSpecialChars checks for special characters.
func ContainsSpecialChars(s string) bool {
	isStringAlphabetic := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString

	return !isStringAlphabetic(s)
}
