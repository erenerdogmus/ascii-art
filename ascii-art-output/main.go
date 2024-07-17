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
		Print(os.Args[1], "standard", false)
	} else if len(os.Args) == 3 {
		Print(os.Args[1], os.Args[2], false)
	} else if len(os.Args) == 4 {
		Print(os.Args[2], os.Args[3], true)
		defer fmt.Println("file created successfully")
	} else {
		fmt.Println("Usage : go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING] [BANNER]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING]")
	}
}

func Print(input string, bannerName string, isEqual bool) {
	if len(input) == 0 {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	var outputFlag *string
	var data *os.File
	var err error
	if isEqual {
		outputFlag = flag.String("output", "", "Output file name")
		flag.Parse()
		if *outputFlag == "" {
			fmt.Println("Usage: --output=<output file")
			return
		}
		data, err = os.Create(*outputFlag)
		if err != nil {
			fmt.Println("file could not be created")
			return
		}
		defer data.Close()
	}
	wordsList := strings.Split(input, "\\n")
	for f, char := range wordsList {
		if char != "" {
			for i, word := range GetAscii(char, bannerName) {
				if isEqual {
					if len(word) > 0 {
						_, err := data.WriteString(word + "\n")
						if err != nil {
							fmt.Println("Error writing to file")
							return
						}
					}
				} else {
					if i < 8 {
						fmt.Println(word)
					}
				}
			}
		} else {
			if isEqual {
				_, err := data.WriteString("\n")
				if err != nil {
					fmt.Println("Error adding a line to a file")
					return
				}
			} else {
				fmt.Println()
			}
		}
		if len(os.Args) == 4 && len(wordsList)-1 == f {
			_, err = data.WriteString("\n")
			if err != nil {
				fmt.Println("Error adding a line to a file")
				return
			}
		}
		if (len(os.Args) == 2 || len(os.Args) == 3) && len(wordsList)-1 == f {
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
