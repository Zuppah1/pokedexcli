package main

import (
	"strings"
)

func cleanInput(text string) []string {

	wordList := strings.Fields(strings.ToLower(text))

	return wordList
}
