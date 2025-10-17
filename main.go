package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Hangman struct {
	secretWord       string
	guesses          []byte
	correctGuesses   []byte
	remainingChances uint
}

func newGame(secretWord string) Hangman {

	return Hangman{
		secretWord:       secretWord,
		guesses:          []byte{},
		correctGuesses:   []byte{},
		remainingChances: 7,
	}
}

func containsPunctuation(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return true
		}
	}
	return false
}

func getSecretWord(fileName string) string {

	var allowedWords []string

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error in %v cause of %v", fileName, err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && len(word) >= 6 && !containsPunctuation(word) {
			allowedWords = append(allowedWords, word)
		}
	}
	randomNum := rand.Intn(len(allowedWords))
	return allowedWords[randomNum]
}

func checkGuess(state Hangman, guess byte) Hangman {

	if state.remainingChances > 1 && strings.ContainsRune(state.secretWord, rune(guess)) && !bytes.Contains(state.guesses, []byte{guess}) {

		state.correctGuesses = append(state.correctGuesses, guess)
		state.guesses = append(state.guesses, guess)

	}
	return state
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
