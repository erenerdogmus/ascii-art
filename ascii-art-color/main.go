package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		print(os.Args[1], false)
	}
	if len(os.Args) == 4 {
		print(os.Args[3], true)
	} else if len(os.Args) == 3 {
		print(os.Args[2], false)
	} else {
		fmt.Println("Usage : go run . [OPTION] [STRING] [STRING]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [OPTION] [STRING]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING]")
	}
}

func print(input string, isEqual bool) {
	if len(input) == 0 {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	var colorFlag *string
	var words string
	if isEqual {
		colorFlag = flag.String("color", "", "set text color")
		flag.Parse()
		if *colorFlag == "" {
			fmt.Println("Usage : go run . --usage=<Color Name>")
			os.Exit(1)
		}
		words = flag.Arg(0)
	} else {
		colorFlag = flag.String("color", "white", "set text color")
		flag.Parse()
		words = input
	}
	wordsList := strings.Split(input, "\\n")
	for _, char := range wordsList {
		if char != "" {
			for i, word := range GetAscii(char, colorFlag, words) {
				if i < 8 {
					fmt.Println(word)
				}
			}
		} else {
			fmt.Println()
		}
	}
}

func GetAscii(asciStr string, colorFlag *string, words string) []string {
	lines := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		asciiValue := (((int(asciStr[i])) - 32) * 9) + 2
		values := LinesRead("standard.txt", asciiValue, asciiValue+8, Control(string(asciStr[i]), words), colorFlag)
		for index, value := range values {
			lines[index] += value
		}
	}
	return lines
}

func Control(inputText string, words string) bool {
	for _, r := range inputText {
		if strings.ContainsRune(words, r) {
			return true
		}
	}
	return false
}

func LinesRead(fileName string, beginning int, finish int, İsEqual bool, colorFlag *string) []string {
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
			if İsEqual {
				lines = append(lines, colors[*colorFlag]+scanner.Text())
			} else {
				lines = append(lines, colors["white"]+scanner.Text())
			}
		}
	}
	return lines
}

var colors = map[string]string{
	"black":          "\033[30m",                // Black
	"red":            "\033[31m",                // Red
	"green":          "\033[32m",                // Green
	"yellow":         "\033[33m",                // Yellow
	"blue":           "\033[34m",                // Blue
	"magenta":        "\033[35m",                // Magenta
	"cyan":           "\033[36m",                // Cyan
	"white":          "\033[37m",                // White
	"bright_black":   "\033[90m",                // Bright Black (Gray)
	"bright_red":     "\033[91m",                // Bright Red
	"bright_green":   "\033[92m",                // Bright Green
	"bright_yellow":  "\033[93m",                // Bright Yellow
	"bright_blue":    "\033[94m",                // Bright Blue
	"bright_magenta": "\033[95m",                // Bright Magenta
	"bright_cyan":    "\033[96m",                // Bright Cyan
	"bright_white":   "\033[97m",                // Bright White
	"orange":         "\u001b[38;2;255;160;16m", // orange
}
