package functions

import (
	"fmt"
	"os"
	"regexp"
)

type Options struct {
	Input      string
	OutputFile string
	Output     string
	Color      string
	ToColor    string
	Banner     string
}

var OptionsData Options

func ArgsChecker(args []string) {
	if len(args) < 2 || len(args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(1)
	}

	colorPattern := regexp.MustCompile(`^--color=`)
	colorCheck := regexp.MustCompile(`^(?:--color=)(\w+)`)

	outputPattern := regexp.MustCompile(`^--output=`)
	outputCheck := regexp.MustCompile(`^(?:--output=)(\w+\W*\w*\.txt)$`)

	if len(args) == 2 {
		OptionsData.Input = args[1]
		fmt.Println(OptionsData)
		return
	}

	if match := colorPattern.FindStringSubmatch(args[1]); match != nil {
		if len(args) < 4 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(1)
		}

		color := colorCheck.FindStringSubmatch(args[1])
		if color == nil {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(1)
		}

		OptionsData.Color = color[1]
		OptionsData.ToColor = args[2]
		OptionsData.Input = args[3]

		if len(args) == 5 {
			OptionsData.Banner = args[4] + ".txt"
		} else {
			OptionsData.Banner = "standard"
		}
	} else if match := outputPattern.FindStringSubmatch(args[1]); match != nil {
		if len(args) < 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --output=<filename.txt> \"something\"")
			os.Exit(1)
		}

		output := outputCheck.FindStringSubmatch(args[1])
		if output == nil {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --output=<filename.txt> \"something\"")
			os.Exit(1)
		}

		OptionsData.OutputFile = match[1]
		OptionsData.Input = args[2]

		if len(args) == 4 {
			OptionsData.Banner = args[3]
		} else {
			OptionsData.Banner = "standard"
		}
	} else {
		OptionsData.Input = args[1]
		OptionsData.Banner = args[2]
	}
}
