package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 4 {
		Print(os.Args[2], os.Args[3], true)
	} else if len(os.Args) == 3 {
		Print(os.Args[1], os.Args[2], false)
	} else {
		fmt.Println("Usage : go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING] [BANNER]")
	}
}

func Print(input string, bannerName string, isEqual bool) {
	if len(input) == 0 {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	var outputflag *string
	if isEqual {
		outputflag = flag.String("align", "", "alignment flag")
		flag.Parse()
		if *outputflag == "" {
			fmt.Println("Usage : go run . --align=<align name:")
			os.Exit(1)
		}
	} else {
		outputflag = flag.String("align", "left", "alignment flag")
		flag.Parse()
	}
	terminalWidth, err := getTerminalWidth()
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	wordsList := strings.Split(input, "\\n")
	for _, char := range wordsList {
		if char != "" {
			for i, word := range GetAscii(char, bannerName) {
				if i < 8 {
					if strings.Contains(*outputflag, "center") {
						printCenter(word, terminalWidth)
					} else if strings.Contains(*outputflag, "left") {
						printLeft(word)
					} else if strings.Contains(*outputflag, "right") {
						printRight(word, terminalWidth)
					} else if strings.Contains(*outputflag, "justify") {
						if count == 0 {
							Printjustify(input, bannerName, outputflag, terminalWidth)
							count++
						}
					} else {
						fmt.Println("Geçersiz hizalama tarzı")
						return
					}
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

func AralikHesapla(kelime string, filename string) []string {
	var satir []string
	satirlar := make([]string, 9)
	for _, harf := range kelime {
		asciidegeri := int(harf)
		asciidegeri = (asciidegeri-32)*9 + 2
		satir = LinesRead(filename, asciidegeri, asciidegeri+8)
		for index, value := range satir {
			satirlar[index] += value
		}
	}
	return satirlar
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

func getTerminalWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	size := strings.TrimSpace(string(out))
	width, err := strconv.Atoi(size)
	if err != nil {
		return 0, err
	}
	return width, nil
}

func printCenter(text string, width int) {
	totalSpaces := width - len(text)
	leftSpaces := totalSpaces / 2
	rightSpaces := totalSpaces - leftSpaces
	fmt.Println(strings.Repeat(" ", leftSpaces) + text + strings.Repeat(" ", rightSpaces))
}

func printLeft(text string) {
	fmt.Println(text)
}

func printRight(text string, width int) {
	totalSpaces := width - len(text)
	fmt.Println(strings.Repeat(" ", totalSpaces) + text)
}

func Printjustify(input string, banner string, align *string, terminal int) {
	kelime := strings.Split(input, " ")
	kelime_sayisi := len(kelime)
	var value [][]string
	result := make([]string, 9)
	for _, i := range kelime {
		value = append(value, AralikHesapla(i, Banner(banner)))
	}
	if kelime_sayisi == 1 {
		for index := range value[0] {
			printCenter(value[0][index], terminal)
		}
		return
	}
	space := 0
	for i := 0; i < len(value); i++ {
		space += len(value[i][0])
	}
	space = (terminal - space) / (kelime_sayisi - 1)
	for j := 0; j < kelime_sayisi; j++ {
		for i := 0; i < len(value[0]); i++ {
			result[i] += value[j][i]
			if j != kelime_sayisi-1 {
				result[i] += strings.Repeat(" ", space)
			}
		}
	}
	for index, i := range result {
		if len(result)-1 != index {
			fmt.Println(i)
		}
	}
}
