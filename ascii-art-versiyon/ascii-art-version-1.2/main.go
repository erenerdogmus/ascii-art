package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var outputFlag *string

func main() {
	if len(os.Args) == 2 {
		Print(os.Args[1], "standard", false)
	} else if len(os.Args) == 3 {
		Print(os.Args[1], os.Args[2], false)
	} else if len(os.Args) == 4 {
		Print(os.Args[2], os.Args[3], true)
	} else {
		fmt.Println("Usage: go run . [STRING]")
		fmt.Println("OR")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("OR")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
}

func Print(text, banner string, isEqual bool) {
	var flagValue *string
	var data *os.File
	var err error
	if isEqual {
		flagValue = FlagControl(os.Args[1])
		data, err = os.Create(*flagValue)
		if err != nil {
			fmt.Println("file could not be created")
			return
		}
		defer data.Close()
		if len(text) == 0 {
			return
		} else if text == "\\n" {
			fmt.Println()
			return
		}
	}
	textList := strings.Split(text, "\\n")
	for _, char := range textList {
		if char != "" {
			for i, r := range GetAscii(char, banner) {
				if i < 8 {
					if isEqual {
						_, err := data.WriteString(r + "\n")
						if err != nil {
							fmt.Println("Error writing to file")
							return
						}
					} else {
						fmt.Println(r)
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
	}

}

func FlagControl(flagText string) *string {
	if strings.Contains(flagText, "--output") {
		return OutputFlag()
	}
	return nil
}

func OutputFlag() *string {
	outputFlag = flag.String("output", "", "Output file name")
	flag.Parse()
	if *outputFlag == "" {
		log.Fatal("Usage: --output=<output file")
	}
	defer fmt.Println("file created successfully")
	return outputFlag
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
