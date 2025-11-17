package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		cleanWords := cleanInput((userInput))

		if len(cleanWords) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", cleanWords[0])
	}
}

func cleanInput(text string) []string {

	wordList := strings.Fields(strings.ToLower(text))

	return wordList
}
