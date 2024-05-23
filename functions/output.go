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

		partToColorIndex := 0
		for count < 8 {
			for i, letter := range part {
				if OptionsData.Color != "" {
					if OptionsData.ToColor != "" {
						if OptionsData.ToColorIndexes == nil || len(OptionsData.ToColorIndexes) == 0 {
							OptionsData.ErrorMsg = "Can't find the target to color"
							return
						}

						if OptionsData.ToColorIndexes[PartIndex] == nil || len(OptionsData.ToColorIndexes[PartIndex]) == 0 {
							continue
						}

						if len(OptionsData.ToColorIndexes[PartIndex]) >= partToColorIndex {
							fmt.Println(len(OptionsData.ToColorIndexes[PartIndex]))
							result += Font[letter][count]
							continue
						}

						if i > OptionsData.ToColorIndexes[PartIndex][partToColorIndex][1] {
							partToColorIndex++
							continue
						}

						if i >= OptionsData.ToColorIndexes[PartIndex][partToColorIndex][0] && i <= OptionsData.ToColorIndexes[PartIndex][partToColorIndex][1] {
							result += OptionsData.Color + Font[letter][count] + Colors["reset"]
						} else {
							result += Font[letter][count]
						}
					} else {
						result += OptionsData.Color + Font[letter][count] + Colors["reset"]
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
