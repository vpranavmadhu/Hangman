package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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
		if word == strings.ToLower(word) && len(word) >= 6 {
			allowedWords = append(allowedWords, word)
		}
	}
	randomNum := rand.Intn(len(allowedWords))
	return allowedWords[randomNum]
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
