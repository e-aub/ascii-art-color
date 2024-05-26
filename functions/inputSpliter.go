package functions

import "strings"

func NewLineIndices() []int {

	arr := []int{}
	for i := 0; i < len(OptionsData.Input)-1; i++ {
		if OptionsData.Input[i:i+2] == "\\n" {
			arr = append(arr, i)
			i++
		}
	}

	if len(arr) == 0 {
		return nil
	}
	return arr
}

func Split() {
	result := strings.Split(OptionsData.Input, "\\n")
	for index, line := range result {
		if line == "" {
			result[index] = "\\n"
		}
	}
	if isSuccessive(result) {
		result = result[:len(result)-1]
	}
	OptionsData.SplicedInput = result
}

func isSuccessive(str []string) bool {
	for _, elem := range str {
		if elem != "\n" {
			return false
		}
	}
	return true
}
