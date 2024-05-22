package functions

import (
	"fmt"
	"os"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	result := ""
	for PartIndex, part := range OptionsData.SplicedInput {
		if part == "\\n" {
			result += "\n"
			continue
		}
		count := 0

		if OptionsData.ErrorMsg != "" {
			break
		}

		for count < 8 {
			for i, letter := range part {
				if OptionsData.Color != "" {
					if OptionsData.ToColor != "" {
						if OptionsData.ToColorIndexes == nil {
							OptionsData.ErrorMsg = "Can't find the sub-string to color"
							// count = 8
							break
						}

						if i >= OptionsData.ToColorIndexes[PartIndex][0] && i <= OptionsData.ToColorIndexes[PartIndex][1] {
							result = result + OptionsData.Color + Font[letter][count] + Colors["reset"]
							continue
						} else {
							result += Font[letter][count]
						}

					} else {
						result = result + OptionsData.Color + Font[letter][count] + Colors["reset"]
					}
				} else {
					result += Font[letter][count]
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
