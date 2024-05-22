package functions

import (
	"sort"
	"strings"
)

func Minimize() {
	str := strings.ReplaceAll(OptionsData.Input, "\n", "")
	var result []rune
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			OptionsData.ErrorMsg = "Invalid Character (" + string(letter) + ")\n\nThe characters must be between ' ' and '~'"
			return
		}
		if !strings.Contains(string(result), string(letter)) {
			result = append(result, letter)
		}
	}

	OptionsData.ToMap = sortRunes(result)
}
//sort the input string
func sortRunes(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}
