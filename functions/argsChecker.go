package functions

import (
	"fmt"
	"regexp"
	"strings"
)

// FlagChecker checks if the provided flags are valid
func FlagChecker() {
	if len(Params.Args) < 1 || len(Params.Args) > 4 {
		UsageErr()
		return
	}
	outputPattern := regexp.MustCompile(`^-{1,2}output`)
	outputCheck := regexp.MustCompile(`^(?:--output=)(.+.txt)$`)
	colorPattern := regexp.MustCompile(`^-{1,2}color`)
	colorCheck := regexp.MustCompile(`^(?:--color=)(\w+)$`)
	validBanner := regexp.MustCompile(`^standard|shadow|enigma|nirvana|standard.txt|shadow.txt|enigma.txt|nirvana.txt$`)
	if len(Params.Args) == 1 {
		if outputPattern.MatchString(Params.Args[0]) || colorPattern.MatchString(Params.Args[0]) {
			UsageErr()
			return
		}
		return
	}
	if outputPattern.MatchString(Params.Args[0]) {
		if len(Params.Args) > 3 {
			UsageErr()
			return
		}
		if output := outputCheck.FindStringSubmatch(Params.Args[0]); output != nil {
			Params.OutputFile = output[1]
			Params.Args = Params.Args[1:]
			return
		}
		UsageErr()
		return
	}
	if colorPattern.MatchString(Params.Args[0]) {
		if color := colorCheck.FindStringSubmatch(Params.Args[0]); color != nil {
			if strings.ToLower(color[1]) == "random" {
				RandomColor()
			}
			colorCode, ok := Colors[strings.ToLower(color[1])]
			if !ok {
				Params.ErrorMsg = "Available colors:\n\n"
				for name, color := range Colors {
					if name == "reset" || color == "" {
						continue
					}
					Params.ErrorMsg += "\033[38;2;" + color + "m" + name + Colors["reset"] + ", "
				}
				return
			}
			Params.Color = "\033[38;2;" + colorCode + "m"
			Params.Args = Params.Args[1:]
			if len(Params.Args) == 2 {
				if !validBanner.MatchString(Params.Args[1]) {
					Params.ToColor = Params.Args[0]
					Params.Args = Params.Args[1:]
				}
				return
			}
			if len(Params.Args) == 1 {
				return
			}
			Params.ToColor = Params.Args[0]
			Params.Args = Params.Args[1:]
			return
		}
		UsageErr()
		return
	}
}

// ArgsChecker checks if the required arguments are provided
func ArgsChecker() {
	if len(Params.Args) == 3 {
		UsageErr()
	} else if len(Params.Args) == 1 {
		Params.Input = Params.Args[0]
		Params.Banner = "standard.txt"
	} else {
		Params.Input = Params.Args[0]
		if !regexp.MustCompile(`\.txt$`).MatchString(Params.Args[1]) {
			Params.Banner = Params.Args[1] + ".txt"
		} else {
			Params.Banner = Params.Args[1]
		}
	}
	fmt.Println("Banner:", Params.Banner, "input:", Params.Input, "toColor:", Params.ToColor)
}

// UsageErr prints the usage error message
func UsageErr() {
	Params.ErrorMsg = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
}
