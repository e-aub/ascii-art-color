package main

import (
	f "ascii-art-color/functions"
	"os"
)

func main() {
	f.ArgsChecker(os.Args)
	if f.OptionsData.Input == "" {
		return
	}
	f.Execute()
}
