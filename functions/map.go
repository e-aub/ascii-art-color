package functions

import (
	"bufio"
	"log"
	"os"
)

var Font map[rune][]string

func MapFont() {
	Font = make(map[rune][]string)
	file, err := os.Open("./templates/" + OptionsData.Banner)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	rIndex := 0
	for scanner.Scan() {
		if lineCount == ((int(OptionsData.ToMap[rIndex])-31)*9)-9 {
			i := 0
			for i < 8 && scanner.Scan() {
				Font[OptionsData.ToMap[rIndex]] = append(Font[OptionsData.ToMap[rIndex]], scanner.Text())
				lineCount++
				i++
			}
			if rIndex != len(OptionsData.ToMap)-1 {
				rIndex++
			}
		}
		lineCount++
	}
}
