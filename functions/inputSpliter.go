package functions

import "strings"

func Split(){
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
		if elem != "\\n" {
			return false
		}
	}
	return true
}
