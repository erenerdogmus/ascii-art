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
		Print(os.Args[1], "standard")
	} else if len(os.Args) == 3 {
		Print(os.Args[1], os.Args[2])
	} else {
		fmt.Println("Usage : go run . [STRING]")
		return
	}
}

func Print(text, banner string) {
	if len(text) == 0 {
		return
	} else if text == "\\n" {
		fmt.Println()
		return
	}
	textList := strings.Split(text, "\\n")
	for _, char := range textList {
		if char != "" {
			for i, r := range GetAscii(char, banner) {
				if i < 8 {
					fmt.Println(r)
				}
			}
		} else {
			fmt.Println()
		}
	}
}

func GetAscii(asciStr, banner string) []string {
	slice := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		assciiValue := ((int(asciStr[i])-32)*9 + 2)
		values := LinesRead(Banner(banner), assciiValue, assciiValue+8)
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

func Banner(bnr string) string {
	if bnr == "standard" || bnr == "standard.txt" {
		return "standard.txt"
	} else if bnr == "shadow" || bnr == "shadow.txt" {
		return "shadow.txt"
	} else if bnr == "thinkertoy" || bnr == "thinkertoy.txt" {
		return "thinkertoy.txt"
	} else {
		return ""
	}
}
