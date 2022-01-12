package main

import (
  "bufio"
  "fmt"
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

	fmt.Println()
	fmt.Println(properTitle(text))
	fmt.Println()
  }

}

func properTitle(input string) string {
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