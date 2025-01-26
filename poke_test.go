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
			input:    " hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hei KNsd JIdi sdfew",
			expected: []string{"hei", "knsd", "jidi", "sdfew"},
		},
	}
	for _, c := range cases {
		actual := cleanInputText(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Not matching %v , %v", word, expectedWord)
			} else {
				fmt.Println("Matching")
			}
		}
	}
}
