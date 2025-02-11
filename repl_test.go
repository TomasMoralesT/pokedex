package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "ALL your bases belong to us",
			expected: []string{"all", "your", "bases", "belong", "to", "us"},
		},
		{
			input:    " hello, world ",
			expected: []string{"hello,", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %v, want %v (lengths differ)", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q) returned %s, want %s (words differ)", c.input, word, expectedWord)
			}
		}
	}
}
