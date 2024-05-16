package functions

import (
	"fmt"
	"log"
	"os"
)

func OutputBuilder() {
	result := ""
	for _, part := range OptionsData.SplicedInput {
		if part == "\\n" {
			result += "\n"
			continue
		}
		count := 0
		for count < 8 {
			for _, letter := range part {
				result += Font[letter][count]
			}
			result += "\n"
			count++
		}
	}
	OptionsData.Output = result
}

func OutputDeliver() {
	if OptionsData.OutputFile == "" {
		fmt.Print(OptionsData.Output)
	} else {
		file, err := os.Create(OptionsData.OutputFile)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		_, err = file.WriteString(OptionsData.Output)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
