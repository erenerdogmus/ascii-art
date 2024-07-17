package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		Print(os.Args[1])
	} else {
		fmt.Println("Usage : go run . [STRING]")
		return
	}
}

func Print(text string) {
	if len(text) == 0 {
		return
	} else if text == "\\n" {
		fmt.Println()
		return
	}
	textList := strings.Split(text, "\\n")
	for _, char := range textList {
		if char != "" {
			for i, r := range GetAscii(char) {
				if i < 8 {
					fmt.Println(r)
				}
			}
		} else {
			fmt.Println()
		}
	}

}

func GetAscii(asciStr string) []string {
	slice := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		assciiValue := (((int(asciStr[i]))-32)*9 + 2)
		values := LinesRead("standard.txt", assciiValue, assciiValue+8)
		for index, value := range values {
			slice[index] += value
		}
	}
	return slice
}

func LinesRead(fileName string, beginning, finish int) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File Not Found")
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
