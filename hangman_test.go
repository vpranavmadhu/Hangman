package main

import (
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
