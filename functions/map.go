package functions

import (
	"bufio"
	"errors"
	"os"
)

var Font map[rune][]string

func MapFont() {
	Font = make(map[rune][]string)
	file, err := os.Open("./templates/" + Params.Banner)
	if err != nil {
		Params.Err = errors.New("fs")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	rIndex := 0
	for scanner.Scan() {
		if lineCount == ((int(Params.ToMap[rIndex])-31)*9)-9 {
			i := 0
			for i < 8 && scanner.Scan() {
				Font[Params.ToMap[rIndex]] = append(Font[Params.ToMap[rIndex]], scanner.Text())
				lineCount++
				i++
			}
			if rIndex != len(Params.ToMap)-1 {
				rIndex++
			}
		}
		lineCount++
	}
}
