package functions

import (
	"errors"
	"regexp"
	"strings"
)

// FlagChecker checks if the provided flags are valid
func FlagChecker() {
	if len(Params.Args) < 1 || len(Params.Args) > 4 {
		Params.Err = errors.New("color")
		return
	}
	if len(Params.Args) == 1 {
		if OutputPattern.MatchString(Params.Args[0]) {
			Params.Err = errors.New("output")
		} else if ColorPattern.MatchString(Params.Args[0]) {
			Params.Err = errors.New("color")
		}
		return
	}

	if OutputPattern.MatchString(Params.Args[0]) {
		if len(Params.Args) > 3 {
			Params.Err = errors.New("output")
			return
		}
		if output := OutputCheck.FindStringSubmatch(Params.Args[0]); output != nil {
			Params.OutputFile = output[1]
			Params.Args = Params.Args[1:]
			return
		}
		Params.Err = errors.New("output")
		return
	}

	if ColorPattern.MatchString(Params.Args[0]) {
		if color := ColorCheck.FindStringSubmatch(Params.Args[0]); color != nil {
			if strings.ToLower(color[1]) == "random" {
				RandomColor()
			}
			if RgbPattern.MatchString(color[1]) {
				RGB(color[1])
			} else if HexPattern.MatchString(color[1]) {
				HexToRgb(color[1])
			} else {
				colorCode, ok := Colors[strings.ToLower(color[1])]
				if !ok {
					Params.Err = errors.New("invalidColor")
					return
				}
				Params.Color = "\033[38;2;" + colorCode + "m"
			}

			Params.Args = Params.Args[1:]

			if len(Params.Args) == 2 {
				if !ValidBanner.MatchString(Params.Args[1]) {
					Params.ToColor = Params.Args[0]
					Params.Args = Params.Args[1:]
					return
				}
				Params.ToColor = Params.Args[0]
				return
			}

			if len(Params.Args) == 1 {
				Params.ToColor = Params.Args[0]
				return
			}

			Params.ToColor = Params.Args[0]
			Params.Args = Params.Args[1:]
			return
		}
		Params.Err = errors.New("color")
		return
	}
}

// ArgsChecker checks if the required arguments are provided
func ArgsChecker() {
	if len(Params.Args) == 3 {
		Params.Err = errors.New("color")
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
}
