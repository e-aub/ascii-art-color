package functions

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// OutputBuilder builds the output string
func OutputBuilder() {
	var result strings.Builder
	tracker := 0
	for _, part := range Params.SplicedInput {
		if part == "\\n" {
			result.WriteString("\n")
			tracker += 2
			continue
		}
		count := 0
		for count < 8 {
			for i, letter := range part {
				currentIndex := i + tracker
				if InRange(currentIndex) {
					result.WriteString(Params.Color + Font[letter][count] + Colors["reset"])
				} else {
					result.WriteString(Font[letter][count])
				}

			}
			result.WriteString("\n")
			count++
		}
		tracker += len(part) + 2
	}
	Params.Output = result.String()
}

// OutputDeliver delivers the output to the console
func OutputDeliver() {
	if Params.OutputFile == "" {
		fmt.Print(Params.Output)
	} else {
		file, err := os.Create(Params.OutputFile)
		if err != nil {
			Params.Err = errors.New("internal")
			return
		}
		defer file.Close()
		_, err = file.WriteString(Params.Output)
		if err != nil {
			Params.Err = errors.New("internal")
			return
		}
	}
}
