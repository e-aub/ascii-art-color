package functions

import (
	"fmt"
	"os"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	result := ""
	index := 0
	for _, part := range OptionsData.SplicedInput {
		if OptionsData.ErrorMsg != "" {
			break
		}

		if part == "\\n" {
			result += "\n"
			continue
		}
		count := 0
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

				if index+1 > len(OptionsData.ToColorIndexes) {
					result += Font[letter][count]
					continue
				}

				first := OptionsData.ToColorIndexes[index]
				last := OptionsData.ToColorIndexes[index+1]

				if i >= first && i <= last {
					result += OptionsData.Color + Font[letter][count]

					if i == last {
						result += Colors["reset"]
					}
				} else {
					result += Font[letter][count]
				}

				if i == last && count == 7 {
					index += 2
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
