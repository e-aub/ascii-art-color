package functions

import (
	"fmt"
	"os"
	"strings"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	result := []string{}
	tempStr := ""
	newLine := -1
	for j := 0; j < 8; j++ {
		PartToColor := 0
		for i := 0; i < len(OptionsData.Input); i++ {
			if i == newLine {
				result = append(result, tempStr)
				tempStr = ""
				i += 1
				continue
			}

			if i < len(OptionsData.Input)-1 && OptionsData.Input[i:i+2] == "\\n" {
				result = append(result, tempStr)
				tempStr = ""
				newLine = i
				i += 1
				continue
			}

			if OptionsData.Color == "" {
				tempStr += Font[rune(OptionsData.Input[i])][j]
				continue
			}

			if OptionsData.ToColorIndexes == nil || len(OptionsData.ToColorIndexes) == 0  {
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
		tempStr += "\n"
	}
	
	OptionsData.Output = strings.Join(result, "") + "\n"
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
