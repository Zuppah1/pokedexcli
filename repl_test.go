package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " random pokemon names ",
			expected: []string{"random", "pokemon", "names"},
		},
		{
			input:    " Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice length incorrect")
		}

		for i := range actual {
			word := actual[i]
			fmt.Println(word)
			expectedWord := c.expected[i]

			if expectedWord != word {
				t.Errorf("words do not match to expected values")
			}
		}
	}
}
