package functions

import (
	"errors"
	"sort"
	"strings"
)

func Minimize() {
	str := strings.ReplaceAll(Params.Input, "\n", "")
	var result []rune
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			Params.Err = errors.New("char")
			return
		}
		if !strings.Contains(string(result), string(letter)) {
			result = append(result, letter)
		}
	}

	Params.ToMap = sortRunes(result)
}

// sort the input string
func sortRunes(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}
