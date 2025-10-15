package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func getSecretWord(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error in %v cause of %v", file, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordList []string
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	randomNumber := rand.Intn(len(wordList))
	fmt.Println("random number:", randomNumber)

	randomWord := wordList[randomNumber]

	return randomWord
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
