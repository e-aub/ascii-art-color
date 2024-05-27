package functions

import (
	"fmt"
	"os"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	result := ""
	tracker := 0
	j := 0

	for _, part := range OptionsData.SplicedInput {
		if part == "\\n" {
			result += "\n"
			tracker += 2
			continue
		}
		count := 0
		for count < 8 {
			j = 0
			for i, letter := range part {
				currentIndex := i + tracker
				if inRange(currentIndex) {
					result = result + OptionsData.Color + Font[letter][count] + "\033[0m"
				} else {
					result += Font[letter][count]
				}
				if j < len(OptionsData.ToColorIndexes) && currentIndex == OptionsData.ToColorIndexes[j][1] {
					j++
				}
			}
			result += "\n"
			count++
		}

		tracker += len(part) + 2
	}

	OptionsData.Output = result
}

func inRange(index int) bool {
	for _, pair := range OptionsData.ToColorIndexes {
		if index >= pair[0] && index <= pair[1] {
			return true
		}
	}
	return false
}

// OutputDeliver delivers the output to the console
func OutputDeliver() {
	if OptionsData.OutputFile == "" {
		fmt.Print(OptionsData.Output)
	} else {
		file, err := os.Create(OptionsData.OutputFile)
		if err != nil {
			OptionsData.ErrorMsg = err.Error()
			return
		}
		defer file.Close()
		_, err = file.WriteString(OptionsData.Output)
		if err != nil {
			OptionsData.ErrorMsg = err.Error()
			return
		}
	}
}
