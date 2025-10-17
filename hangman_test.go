package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func createDictFile(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "hangman-dict")
	if err != nil {
		fmt.Println("Couldn't create temp file.")
	}
	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

func TestSecretWordNoCapitals(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion", "Elephant", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestSecretWordLength(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion", "pen", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("should get monkey but Got %s", secretWord)
	}

}

func TestSecretWordPunctuations(t *testing.T) {
	wordList, err := createDictFile([]string{"elephant's", "balloon's", "monkey"})
	if err != nil {
		t.Errorf("Couln't create word list. Can't proceed with test: %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get monkey but Got %s", secretWord)
	}
}

func TestCorrectGuess(t *testing.T) {
	secretWord := "soldier"
	guess := 's'
	currentState := newGame(secretWord)
	newState := checkGuess(currentState, byte(guess))

	expected := Hangman{
		secretWord:       currentState.secretWord,
		guesses:          append(currentState.guesses, byte(guess)),
		correctGuesses:   append(currentState.correctGuesses, byte(guess)),
		remainingChances: 7,
	}

	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guesses, expected.guesses) {
		t.Errorf("Guess should be %q but got %q", expected.guesses, newState.guesses)
	}

	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}

	if !(newState.remainingChances == expected.remainingChances) {
		t.Errorf("Remaining chances is modified")
	}
}
