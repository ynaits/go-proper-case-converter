package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/atotto/clipboard"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("======================================================")
	fmt.Println(color.Bold,"PROPER CASE CONVERTER",color.Reset)
	fmt.Println(color.Colorize(color.Green," Build by: Ratheesh Rajan for YNA"))
	fmt.Println(color.Colorize(color.Red," Copyright © 2022 YNA ITS"))
	fmt.Println("")
	fmt.Println(" How to use:")
	fmt.Println(" -----------")
	fmt.Println(" Enter or paste word or words, press ENTER key.")
	fmt.Println(color.Colorize(color.Yellow," Converted text will copied to clipboard automatically!"))
	fmt.Println("======================================================")
	fmt.Println("")
	for {
		fmt.Print("Enter texts to Convert :=> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		converted := properCase(text)
		fmt.Println(color.Colorize(color.Green, converted))

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
