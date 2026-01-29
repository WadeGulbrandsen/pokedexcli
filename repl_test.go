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
			input:    "   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur  PIKACHU   ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of %v doesn't match length of expected %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word %d expected %s but got %s", i+1, expectedWord, word)
			}
		}
	}
}
