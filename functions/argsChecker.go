package functions

import (
	"fmt"
	"regexp"
	"strings"
)

// Options struct contains the options for the program
type Options struct {
	args           []string
	Input          string
	SplicedInput   []string
	ToMap          []rune
	OutputFile     string
	Output         string
	Color          string
	ToColor        string
	ToColorIndexes []int
	Banner         string
	ErrorMsg       string
}

var OptionsData Options

// FlagChecker checks if the provided flags are valid
func FlagChecker(args []string) {
	if len(args) < 1 || len(args) > 4 {
		UsageErr()
		return
	}

	outputPattern := regexp.MustCompile(`^-{1,2}output`)
	outputCheck := regexp.MustCompile(`^(?:--output=)(.+.txt)$`)
	colorPattern := regexp.MustCompile(`^-{1,2}color`)
	colorCheck := regexp.MustCompile(`^(?:--color=)(\w+)$`)

	validBanner := regexp.MustCompile(`^standard|shadow|enigma|nirvana$`)

	if len(args) == 1 {
		if outputPattern.MatchString(args[0]) {
			UsageErr()
			return
		}

		OptionsData.args = args
		return
	}

	if outputPattern.MatchString(args[0]) {
		if len(args) > 3 {
			UsageErr()
			return
		}
		if output := outputCheck.FindStringSubmatch(args[0]); output != nil {
			OptionsData.OutputFile = output[1]
			OptionsData.args = args[1:]
			return
		}

		UsageErr()
		return
	}

	if colorPattern.MatchString(args[0]) {
		if color := colorCheck.FindStringSubmatch(args[0]); color != nil {
			if strings.ToLower(color[1]) == "random" {
				RandomColor()
			}

			colorCode, ok := Colors[strings.ToLower(color[1])]
			if !ok {
				OptionsData.ErrorMsg = "Available colors:\n\n"
				for name, color := range Colors {
					if name == "reset" || color == "" {
						continue
					}
					OptionsData.ErrorMsg += "\033[38;2;" + color + "m" + name + Colors["reset"] + ", "
				}
				return
			}

			OptionsData.Color = "\033[38;2;" + colorCode + "m"

			args = args[1:]

			if len(args) == 2 {
				if validBanner.MatchString(args[1]) {
					OptionsData.args = args
				} else {
					OptionsData.ToColor = args[0]
					OptionsData.args = args[1:]
				}
				return
			}

			if len(args) == 1 {
				OptionsData.args = args
				return
			}

			OptionsData.ToColor = args[0]
			OptionsData.args = args[1:]
			return
		}

		UsageErr()
		return
	}

	OptionsData.args = args
}

// ArgsChecker checks if the required arguments are provided
func ArgsChecker() {
	if len(OptionsData.args) == 3 {
		UsageErr()
	} else if len(OptionsData.args) == 1 {
		OptionsData.Input = OptionsData.args[0]
		OptionsData.Banner = "standard.txt"
	} else {
		OptionsData.Input = OptionsData.args[0]
		if !regexp.MustCompile(`\.txt$`).MatchString(OptionsData.args[1]) {
			OptionsData.Banner = OptionsData.args[1] + ".txt"
		} else {
			OptionsData.Banner = OptionsData.args[1]
		}
	}
	fmt.Println("Banner:", OptionsData.Banner, "input:", OptionsData.Input, "toColor:", OptionsData.ToColor)
}

// UsageErr prints the usage error message
func UsageErr() {
	OptionsData.ErrorMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
}
