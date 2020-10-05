package rstring_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/patil-ashutosh/go-regex-utility/regex/rstring"
)

func TestCountStringOccurrenceInString(t *testing.T) {
	type test struct {
		name               string
		input              string
		subStr             string
		expectedOccurrence int
	}

	tests := []test{
		{name: "HaveOccurrence", input: "this is string", subStr: "is", expectedOccurrence: 2},
		{name: "NoOccurrence", input: "this is string", subStr: "test", expectedOccurrence: 0},
		{name: "NoOccurrenceEmptyString", input: "", subStr: "is", expectedOccurrence: 0},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			occurrences := rstring.CountStringOccurrenceInString(tc.input, tc.subStr)
			if occurrences != tc.expectedOccurrence {
				t.Errorf("CountStringOccurrenceInString(%v,%v) Failed: expected %d, actual %d",
					tc.subStr, tc.input, tc.expectedOccurrence, occurrences)
			}
		})
	}
}

func TestSplitString(t *testing.T) {
	tests := map[string]struct {
		input    string
		sep      string
		n        int
		expected []string
	}{
		"basic":                 {input: "a,b,c", sep: ",", n: -1, expected: []string{"a", "b", "c"}},
		"wrong sep":             {input: "a/b/c", sep: ",", n: -1, expected: []string{"a/b/c"}},
		"no sep in string":      {input: "abc", sep: "/", n: -1, expected: []string{"abc"}},
		"trailing sep":          {input: "a/b/c/", sep: "/", n: -1, expected: []string{"a", "b", "c", ""}},
		"sep with max split 2":  {input: "a.b.c.d.e", sep: "\\.", n: 2, expected: []string{"a", "b.c.d.e"}},
		"no separator in input": {input: "mango is fruit", sep: "\\s+", n: -1, expected: []string{"mango", "is", "fruit"}},
	}

	for name, tc := range tests {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			result := rstring.SplitString(tc.n, tc.input, tc.sep)
			diff := reflect.DeepEqual(tc.expected, result)
			if !diff {
				t.Errorf("SplitString(%d, %v, %v) Failed: expected %v, actual %v",
					tc.n, tc.input, tc.sep, tc.expected, result)
			}
		})
	}
}

func TestRemoveNonAlphaNumeric(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"with-dashes":      {input: "enc-yclopedia-", expected: "encyclopedia"},
		"with-underscores": {input: "encyclop_edia__", expected: "encyclopedia"},
		"mixed-characters": {input: "encyc@lo^pe%dia-_", expected: "encyclopedia"},
	}

	for name, tc := range tests {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			result := rstring.RemoveNonAlphaNumeric(tc.input)
			diff := reflect.DeepEqual(tc.expected, result)
			if !diff {
				t.Errorf("RemoveNonAlphaNumeric(%v) Failed: expected %v, got %v",
					tc.input, tc.expected, result)
			}
		})
	}
}

func ExampleCountStringOccurrenceInString() {
	fmt.Println(rstring.CountStringOccurrenceInString("this is string", "is"))
	// Output: 2
}

func ExampleSplitString() {
	fmt.Println(rstring.SplitString(-1, "a,b,c", ","))
	// Output: [a b c]
}

func ExampleSplitString_other() {
	fmt.Println(rstring.SplitString(-1, "mango     is     fruit"))
	// Output: [mango is fruit]
}
