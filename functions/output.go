package functions

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	result := []string{}
	newLines := NewLineIndices()
	tempStr := ""
	for j := 0; j < 8; j++ {
		PartToColor := 0
		k := 0

		for i := 0; i < len(OptionsData.Input); i++ {
			if slices.Contains(newLines, i) {
				if k > len(result)-1 {
					if tempStr != "" {
						result = append(result, tempStr+"\n")
					}
					if j == 7 {
						result = append(result, "\n")
					}
				} else {
					if tempStr != "\n" && tempStr != "" {
						result[k] += tempStr + "\n"
					}
				}
				tempStr = ""
				i += 1
				k++
				continue
			}

			if OptionsData.Color == "" {
				tempStr += Font[rune(OptionsData.Input[i])][j]
				continue
			}

			if OptionsData.ToColorIndexes == nil || len(OptionsData.ToColorIndexes) == 0 {
				tempStr += OptionsData.Color + Font[rune(OptionsData.Input[i])][j] + Colors["reset"]
				continue
			}

			if PartToColor >= len(OptionsData.ToColorIndexes) {
				tempStr += Font[rune(OptionsData.Input[i])][j]
				continue
			}

			first := OptionsData.ToColorIndexes[PartToColor][0]
			last := OptionsData.ToColorIndexes[PartToColor][1]

			if i >= first && i <= last {
				tempStr += OptionsData.Color + Font[rune(OptionsData.Input[i])][j] + Colors["reset"]
			} else {
				tempStr += Font[rune(OptionsData.Input[i])][j]
			}
			if i == last {
				PartToColor++
			}
		}
		if tempStr != "" {
			tempStr += "\n"
		}

		if k > len(result)-1 {
			if tempStr != "" {
				result = append(result, tempStr)
			}

			if j == 7 {
				result = append(result, "\n")
			}
		} else {
			if tempStr != "" {
				result[k] += tempStr
			}

		}
		tempStr = ""
	}
	
	if isSuccessive(result) {
		result = result[:len(result)-1]
	}

	OptionsData.Output = strings.Join(result, "")
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
