package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		Print(os.Args[1])
	} else {
		fmt.Println("Usage : go run . [STRING]")
		return
	}
}

func Print(input string) {
	if len(input) == 0 {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	wordsList := strings.Split(input, "\\n")
	for _, char := range wordsList {
		if char != "" {
			for i, word := range GetAscii(char) {
				if i < 8 {
					fmt.Println(word)
				}
			}
		} else {
			fmt.Println()
		}
	}
}

func GetAscii(asciStr string) []string {
	lines := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		assciiValue := (((int(asciStr[i]))-32)*9 + 2)
		values := LinesRead("standard.txt", assciiValue, assciiValue+8)
		for index, value := range values {
			lines[index] += value
		}
	}
	return lines
}

func LinesRead(fileName string, beginning int, finish int) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	numberOfLines := 0
	for scanner.Scan() {
		numberOfLines++
		if numberOfLines >= beginning && numberOfLines <= finish {
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}
