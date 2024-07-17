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

var colorFlag *string
var outputFlag *string
var alignflag *string

func main() {
	if len(os.Args) == 4 {
		if strings.Contains(os.Args[1], "color") {
			input := os.Args[3]
			bannerName := "standard"
			isEqual := true
			Print(input, bannerName, isEqual)
		}
		if strings.Contains(os.Args[1], "output") || strings.Contains(os.Args[1], "align") {
			input := os.Args[2]
			bannerName := os.Args[3]
			isEqual := true
			Print(input, bannerName, isEqual)
			if strings.Contains(os.Args[1], "output") {
				fmt.Println("file created successfully")
			}
		}
	} else if len(os.Args) == 3 {
		if strings.Contains(os.Args[1], "color") {
			input := os.Args[2]
			bannerName := "standard"
			isEqual := true
			Print(input, bannerName, isEqual)
		} else {
			input := os.Args[1]
			bannerName := os.Args[2]
			isEqual := false
			Print(input, bannerName, isEqual)
		}
	} else if len(os.Args) == 2 {
		input := os.Args[1]
		bannerName := "standard"
		isEqual := false
		Print(input, bannerName, isEqual)
	} else {
		fmt.Println("Usage : go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING] [BANNER]")
		fmt.Println("Or")
		fmt.Println("Usage : go run . [STRING]")
	}
}

func FlagControl(flagText string) *string {
	if strings.Contains(flagText, "--color") {
		colorFlag = flag.String("color", "", "set text color")
		flag.Parse()
		if *colorFlag == "" {
			fmt.Println("Usage : go run . --usage=<Color Name>")
			os.Exit(1)
		}
		return colorFlag
	} else if strings.Contains(flagText, "--output") {
		outputFlag = flag.String("output", "", "Output file name")
		flag.Parse()
		if *outputFlag == "" {
			fmt.Println("Usage: --output=<output file")
			os.Exit(1)
		}
		return outputFlag
	} else if strings.Contains(flagText, "--align") {
		alignflag = flag.String("align", "", "alignment flag")
		flag.Parse()
		if *alignflag == "" {
			fmt.Println("Usage : go run . --align=<align name:")
			os.Exit(1)
		}
		return alignflag
	}
	return nil
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

func Print(input string, bannerName string, isEqual bool) {
	if len(input) == 0 {
		return
	} else if input == "\\n" {
		fmt.Println()
		return
	}
	var Flag *string
	var words string
	var data *os.File
	var err error
	if isEqual {
		FlagPtr := FlagControl(os.Args[1])
		if FlagPtr != nil {
			Flag = FlagPtr
		} else {
			fmt.Println("Flag is nil")
		}
		if strings.Contains(os.Args[1], "color") {
			words = flag.Arg(0)
		}
		if strings.Contains(os.Args[1], "output") {
			data, err = os.Create(*Flag)
			if err != nil {
				fmt.Println("file could not be created")
				os.Exit(1)
			}
			defer data.Close()
		}
	} else {
		if strings.Contains(os.Args[1], "color") {
			Flag = flag.String("color", "white", "set text color")
			flag.Parse()
			words = input
		} else if strings.Contains(os.Args[1], "align") {
			Flag = flag.String("align", "left", "alignment flag")
			flag.Parse()
		}
	}
	terminalWidth, err := getTerminalWidth()
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	wordsList := strings.Split(input, "\\n")
	if strings.Contains(os.Args[1], "align") {
		for _, char := range wordsList {
			if char != "" {
				for i, word := range GetAsciiAlign(char, bannerName) {
					if i < 8 {
						if strings.Contains(os.Args[1], "center") {
							printCenter(word, terminalWidth)
						} else if strings.Contains(os.Args[1], "left") {
							printLeft(word)
						} else if strings.Contains(os.Args[1], "right") {
							printRight(word, terminalWidth)
						} else if strings.Contains(os.Args[1], "justify") {
							if count == 0 {
								Printjustify(input, bannerName, Flag, terminalWidth)
								count++
							}
						} else {
							fmt.Println("Geçersiz hizalama tarzı")
							return
						}
					} else {
						fmt.Println()
					}
				}
			}
		}
	} else {
		for _, char := range wordsList {
			if char != "" {
				for i, word := range GetAscii(char, Flag, words, bannerName) {
					if isEqual {
						if strings.Contains(os.Args[1], "output") {
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
					} else {
						if i < 8 {
							fmt.Println(word)
						}
					}

				}
			} else {
				if strings.Contains(os.Args[1], "output") {
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

}

func GetAscii(asciStr string, colorFlag *string, words string, bannerName string) []string {
	lines := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		asciiValue := (((int(asciStr[i])) - 32) * 9) + 2
		values := LinesRead(Banner(bannerName), asciiValue, asciiValue+8, Control(string(asciStr[i]), words), colorFlag)
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
				if strings.Contains(os.Args[1], "output") {
					lines = append(lines, scanner.Text())
				} else {
					lines = append(lines, colors["white"]+scanner.Text())
				}
			}
		}
	}
	return lines
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
	lst := []int{}
	lst = append(lst, (len(text)))
	imp := lst[len(lst)-1]
	totalSpaces := width - imp
	fmt.Println(strings.Repeat(" ", totalSpaces) + text)
}

func Printjustify(input string, banner string, align *string, terminal int) {
	kelime := strings.Split(input, " ")
	kelime_sayisi := len(kelime)
	var value [][]string
	result := make([]string, 9)
	if *align == "justify" {
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
		return
	}
}

func AralikHesapla(kelime string, filename string) []string {
	var satir []string
	satirlar := make([]string, 9)
	for _, harf := range kelime {
		asciidegeri := int(harf)
		asciidegeri = (asciidegeri-32)*9 + 2
		satir = SatirOku(filename, asciidegeri, asciidegeri+8)
		for index, value := range satir {
			satirlar[index] += value
		}
	}
	return satirlar
}

func SatirOku(filename string, baslangic, bitis int) []string {
	file, _ := os.Open(filename) // Dosya acilir
	var result []string
	defer file.Close()

	scanner := bufio.NewScanner(file)
	satirsayisi := 0
	for scanner.Scan() { // dosya içeriği okunur
		satirsayisi++
		if satirsayisi >= baslangic && satirsayisi <= bitis { // belirtilen aralıktaki satırlar alınır
			result = append(result, scanner.Text()) // belirtilen aralıktaki satırlar sonuca eklenir
		}
	}
	return result
}

func GetAsciiAlign(asciStr string, bannerName string) []string {
	lines := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {
		assciiValue := (((int(asciStr[i]))-32)*9 + 2)
		values := LinesReadAlign(Banner(bannerName), assciiValue, assciiValue+8)
		for index, value := range values {
			lines[index] += value
		}
	}
	return lines
}

func LinesReadAlign(fileName string, beginning int, finish int) []string {
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
