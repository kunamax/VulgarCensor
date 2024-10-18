package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func checkIfVulgar(sub string, pattern string) bool {
	sub = " " + sub + " "
	if strings.Contains(pattern, strings.ToLower(sub)) && !(sub == "" || sub == " ") {
		return true
	} else {
		return false
	}
}

func main() {
	var path string
	fmt.Scanln(&path)

	pwd, _ := os.Getwd()
	// TODO change before evaluation
	//pwd += "/Obscene Vocabulary Checker/task/"
	pwd += "/"

	dat, err := os.ReadFile(pwd + path)
	content := " "
	content += string(dat)
	content = strings.ReplaceAll(content, "\n", " ")
	content += " "

	if err != nil {
		fmt.Println(errors.New("wrong path"))
		os.Exit(3)
	}

	for {
		inputReader := bufio.NewReader(os.Stdin)
		word, _ := inputReader.ReadString('\n')
		word = strings.ReplaceAll(word, "\n", "")
		if word == "exit" {
			fmt.Println("Bye!")
			os.Exit(3)
		}
		var sentences []string
		sentences = strings.Split(word, " ")

		for i := range sentences {
			if checkIfVulgar(sentences[i], content) {
				for j := 0; j < len(sentences[i]); j++ {
					fmt.Print("*")
				}
				fmt.Print(" ")
			} else {
				fmt.Print(sentences[i])
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
