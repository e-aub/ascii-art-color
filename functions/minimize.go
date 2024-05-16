package functions

import (
	"log"
	"sort"
	"strings"
)

func Minimize(str string) []rune {
	str = strings.ReplaceAll(str, "\n", "")
	var result []rune
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			log.Fatalln("Invalid Input : characters must be between ' ' and '~'")
		}
		if !strings.Contains(string(result), string(letter)) {
			result = append(result, letter)
		}
	}
	return sortRunes(result)
}

func sortRunes(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}
