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

// RemoveLineBreaks removes all line breaks from a string using a regex.
// As per https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html#lineending,
// Any Unicode linebreak sequence, is equivalent to \u000D\u000A|[\u000A\u000B\u000C\u000D\u0085\u2028\u2029] .
func RemoveLineBreaks(s string) string {
	re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000A}\x{000B}\x{000C}\x{000D}\x{0085}\x{2028}\x{2029}]`)

	return re.ReplaceAllString(s, ``)
}
