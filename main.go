package main

import "fmt"

func getSecretWord(wordList string) string {

	return "pranav"
}

func main() {
	fmt.Println(getSecretWord("usr/share/dict/words"))
}
