package functions

import (
	"fmt"
	"os"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	result := ""
	for j, part := range OptionsData.SplicedInput {
		if OptionsData.ErrorMsg != "" {
			break
		}

		if part == "\\n" {
			result += "\n"
			continue
		}
		count := 0
		index := 0
		lastAdded := -1
		for count < 8 {
			if OptionsData.ErrorMsg != "" {
				break
			}
			for i, letter := range part {
				if OptionsData.ErrorMsg != "" {
					break
				}

				if OptionsData.Color == "" {
					result += Font[letter][count]
					continue
				}

				if OptionsData.ToColor == "" {
					result += OptionsData.ToColor + Font[letter][count] + Colors["reset"]
					continue
				}

				if j >= len(OptionsData.ToColorIndexes) || OptionsData.ToColorIndexes[j] == nil || len(OptionsData.ToColorIndexes[j]) == 0 {
					result += Font[letter][count]
					continue
				}

				if index >= len(OptionsData.ToColorIndexes[j]) {
					result += Font[letter][count]
					continue
				}

				for k := 0; k < len(OptionsData.ToColorIndexes[j]); k += 2 {
					if i == lastAdded {
						continue
					}
					if i >= OptionsData.ToColorIndexes[j][k] && i <= OptionsData.ToColorIndexes[j][k+1] {
						result += OptionsData.Color + Font[letter][count] + Colors["reset"]
					} else {
						result += Font[letter][count]
					}

					lastAdded = i
				}
			}

			result += "\n"
			count++
		}
	}
	OptionsData.Output = result
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
