package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var outputFlag *string
var colorFlag *string
var flags []string
var contains []string
var isMultiFlag bool
var empty []string

func main() {
	if len(os.Args) == 2 {
		Control(os.Args[1], "standard", false, false, empty, empty)
	} else if len(os.Args) == 3 {
		if strings.Contains(os.Args[1], "color") {
			Control(os.Args[2], "standard", false, false, empty, empty)
		} else if strings.Contains(os.Args[1], "output") {
			Control(os.Args[2], "standard", true, false, empty, empty)
		} else {
			Control(os.Args[1], os.Args[2], false, false, empty, empty)
		}
	} else if len(os.Args) == 4 {
		flags, contains, isMultiFlag = MultiflagCtrl(os.Args[1:])
		if isMultiFlag {
			fmt.Println(flags, contains)
			Control(os.Args[3], "standard", false, true, flags, contains)
		} else {
			if strings.Contains(os.Args[1], "color") {
				if matched, _ := regexp.MatchString(`--color.*`, os.Args[1]); matched && regexp.MustCompile(`(\.txt|d|w|y)`).MatchString(os.Args[3]) {
					Control(os.Args[2], os.Args[3], false, false, empty, empty)
				} else if strings.Contains(os.Args[1], "color") {
					Control(os.Args[3], "standard", true, false, empty, empty)
				}
			} else if strings.Contains(os.Args[1], "output") {
				Control(os.Args[2], os.Args[3], true, false, empty, empty)
			}
		}
	} else {
		Usage()
		return
	}
}

func Control(text, banner string, isEqual, isMultiFlag bool, flags, contains []string) {
	if len(text) == 0 {
		return
	} else if text == "\\n" {
		fmt.Println()
		return
	}
	var flagValue *string
	var data *os.File
	var err error
	var words string
	if isEqual || (isMultiFlag) || (isEqual && isMultiFlag) {
		flagValue = FlagControl(os.Args[1])
		if isMultiFlag {
			for i := 0; i < len(flags); i++ {

			}
		}
		if len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--output") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--output")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--output"))) {
			if len(os.Args) >= 3 && strings.Contains(os.Args[2], "--output") {
				flagValue = FlagControl(os.Args[2])
			} else if len(os.Args) >= 4 && strings.Contains(os.Args[3], "--output") {
				flagValue = FlagControl(os.Args[3])
			}
			data, err = os.Create(*flagValue)
			if err != nil {
				fmt.Println("file could not be created")
				return
			}
			defer data.Close()
		} else if len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--color") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--color")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--color"))) {
			if len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--color")) {
				words = flag.Arg(0)
			} else if len(os.Args) >= 3 && strings.Contains(os.Args[2], "--color") {
				words = flag.Arg(1)
			} else if len(os.Args) >= 4 && strings.Contains(os.Args[3], "--color") {
				words = flag.Arg(2)
			}
		}
	} else {
		if len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--color") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--color")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--color"))) {
			flagValue = flag.String("color", "white", "set text color")
			flag.Parse()
			words = text
		}
	}

	print(text, banner, isEqual, flagValue, words, data)
}

func print(text string, banner string, isEqual bool, flagValue *string, words string, data *os.File) {
	textList := strings.Split(text, "\\n")
	for _, char := range textList {
		if char != "" {
			for i, r := range GetAscii(char, banner, flagValue, words) {
				if i < 8 {
					if isEqual && len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--output") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--output")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--output"))) {
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
			if isEqual && len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--output") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--output")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--output"))) {
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

func MultiflagCtrl(args []string) ([]string, []string, bool) {
	var firstFlags []string
	isMultiFlag = false
	contains := make([]string, 3)
	if len(args) == 0 || len(args[0]) < 8 {
		fmt.Println("Incorrect usage")
		return flags, contains, isMultiFlag
	}
	for _, argm := range args {
		if strings.Contains(argm, "--") {
			firstFlags = append(firstFlags, argm)
		}
	}
	for index, flag := range firstFlags {
		if flag[:3] == "--o" {
			flags = append(flags, "output")
			contains[index] = flag[9:]
		} else if flag[:3] == "--c" {
			flags = append(flags, "color")
			contains[index] = flag[8:]
		} else {
			fmt.Println("Incorrect flag usage")
		}
	}
	if len(flags) > 1 {
		isMultiFlag = true
	}
	if isMultiFlag {
		for index, char := range args {
			if strings.Contains(char, "--color") || strings.Contains(char, "--output") {
				args = RemoveElement(args, index)
			}
		}
		args = RemoveElement(args, 0)
	}
	return flags, contains, isMultiFlag
}

func RemoveElement(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func FlagControl(flagText string) *string {
	if len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--output") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--output")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--output"))) {
		return OutputFlag()
	} else if len(os.Args) >= 2 && (strings.Contains(os.Args[1], "--color") || (len(os.Args) >= 3 && strings.Contains(os.Args[2], "--color")) || (len(os.Args) >= 4 && strings.Contains(os.Args[3], "--color"))) {
		return ColorFlag()
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

func ColorFlag() *string {
	colorFlag = flag.String("color", "", "set text color")
	flag.Parse()
	if *colorFlag == "" {
		log.Fatal("Usage : go run . --usage=<Color Name>")
	}
	return colorFlag
}

func GetAscii(asciStr, bnr string, colorFlag *string, words string) []string {
	slice := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		assciiValue := ((int(asciStr[i])-32)*9 + 2)
		values := LinesRead(Banner(bnr), assciiValue, assciiValue+8, ControlColor(string(asciStr[i]), words), colorFlag)
		for index, value := range values {
			slice[index] += value
		}
	}
	return slice
}

func LinesRead(fileName string, beginning int, finish int, İsEqual bool, colorFlag *string) []string {
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
			if İsEqual {
				lines = append(lines, colors[*colorFlag]+scanner.Text())
			} else {
				lines = append(lines, colors["white"]+scanner.Text())
			}
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

func ControlColor(inputText string, words string) bool {
	for _, r := range inputText {
		if strings.ContainsRune(words, r) {
			return true
		}
	}
	return false
}

var colors = map[string]string{
	"black":  "\033[30m",                // Black
	"red":    "\033[31m",                // Red
	"green":  "\033[32m",                // Green
	"yellow": "\033[33m",                // Yellow
	"blue":   "\033[34m",                // Blue
	"white":  "\033[37m",                // White
	"orange": "\u001b[38;2;255;160;16m", // orange
}

func Usage() {
	fmt.Println("Usage: go run . [STRING]")
	fmt.Println("OR")
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("OR")
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
}
