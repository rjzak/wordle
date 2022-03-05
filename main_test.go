package main

import (
	"testing"
)

// Ensure that all words are five characters long
func TestWordLength(t *testing.T) {
	for index, word := range wordsList {
		if len(word) != 5 {
			t.Errorf("Word #%d %s is not five characters long!", index, word)
		}
	}
}
