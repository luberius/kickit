package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func processWord(word string) string {
	if string([]rune(word)[0]) == "~" {
		key := trimLeftChar(word)
		return key
	}

	return word
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func main() {
	//rootPath := "C:\\Users\\User\\Project\\KickIt\\blog"

	file, err := ioutil.ReadFile("template/model.kicktpl")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		if strings.Contains(line, "~") {
			words := strings.Fields(line)
			for j, word := range words {
				words[j] = processWord(word)
			}

			lines[i] = strings.Join(words, " ")
		}
	}

	output := strings.Join(lines, "\n")
	fmt.Print(output)
}
