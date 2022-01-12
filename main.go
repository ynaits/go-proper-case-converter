package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Case Converter")
	fmt.Println("=============")

	for {
		fmt.Print("Enter text to Convert :=> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		converted := properCase(text)
		fmt.Println(converted)
		clipboard.WriteAll(converted)
	}

}

func properCase(input string) string {
	toLower := strings.ToLower(input)
	return properTitleCase(toLower)
}

func properTitleCase(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
