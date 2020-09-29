package testing

import (
	"reflect"
	"testing"

	"github.com/patil-ashutosh/RegexUtility/regex/rstring"
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
		t.Run(tc.name, func(t *testing.T) {
			occurrences := rstring.CountStringOccurrenceInString(tc.subStr, tc.input)
			if occurrences != tc.expectedOccurrence {
				t.Errorf("CountStringOccurrenceInString(%v,%v) Failed: expected %d, actual %d",
					tc.subStr, tc.input, tc.expectedOccurrence, occurrences)
			}
		})

	}
}

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input    string
		sep      string
		n        int
		expected []string
	}{
		"basic":                {input: "a,b,c", sep: ",", n: -1, expected: []string{"a", "b", "c"}},
		"wrong sep":            {input: "a/b/c", sep: ",", n: -1, expected: []string{"a/b/c"}},
		"no sep":               {input: "abc", sep: "/", n: -1, expected: []string{"abc"}},
		"trailing sep":         {input: "a/b/c/", sep: "/", n: -1, expected: []string{"a", "b", "c", ""}},
		"sep with max split 2": {input: "a.b.c.d.e", sep: "\\.", n: 2, expected: []string{"a", "b.c.d.e"}},
	}

	for name, tc := range tests {
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
