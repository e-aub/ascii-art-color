package functions

import (
	"bufio"
	"log"
	"os"
)

var Font map[rune][]string

func MapFont(banner string, toMap []rune) {
	Font = make(map[rune][]string)
	file, err := os.Open("./templates/" + banner)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	rIndex := 0
	for scanner.Scan() {
		if lineCount == ((int(toMap[rIndex])-31)*9)-9 {
			i := 0
			for i < 8 && scanner.Scan() {
				Font[toMap[rIndex]] = append(Font[toMap[rIndex]], scanner.Text())
				lineCount++
				i++
			}
			if rIndex != len(toMap)-1 {
				rIndex++
			}
		}
		lineCount++
	}
}
