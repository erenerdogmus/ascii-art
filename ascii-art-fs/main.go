package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		Print(os.Args[1], "standard")
	} else if len(os.Args) == 3 {
		Print(os.Args[1], os.Args[2])
	} else {
		fmt.Println("Usage : go run . [STRING] [BANNER]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING]")
	}
}

func Print(input string, bannerName string) {
	if len(input) == 0 {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	wordsList := strings.Split(input, "\\n")
	for _, char := range wordsList {
		if char != "" {
			for i, word := range GetAscii(char, bannerName) {
				if i < 8 {
					fmt.Println(word)
				}
			}
		} else {
			fmt.Println()
		}
	}
}

func Banner(bannerName string) string {
	if bannerName == "standard" {
		return "standard.txt"
	} else if bannerName == "shadow" {
		return "shadow.txt"
	} else if bannerName == "thinkertoy" {
		return "thinkertoy.txt"
	} else {
		return ""
	}
}

func GetAscii(asciStr string, bannerName string) []string {
	lines := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		assciiValue := (((int(asciStr[i]))-32)*9 + 2)
		values := LinesRead(Banner(bannerName), assciiValue, assciiValue+8)
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
