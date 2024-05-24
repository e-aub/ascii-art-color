package main

import (
	f "ascii-art-output/functions"
	"fmt"
	"os"
)

func main() {
	// Check if the provided flags and args are valid
	f.FlagChecker(os.Args[1:])
	if f.OptionsData.ErrorMsg != "" {
		fmt.Println(f.OptionsData.ErrorMsg)
		return
	}
	f.ArgsChecker()
	if f.OptionsData.ErrorMsg != "" {
		fmt.Println(f.OptionsData.ErrorMsg)
		return
	}

	if f.OptionsData.Input == "" {
		return
	}

	f.ToColorIndexes()

	// Minimize the input string
	f.Minimize()
	if f.OptionsData.ErrorMsg != "" {
		fmt.Println(f.OptionsData.ErrorMsg)
		return
	}
	// Map the input string to the selected font
	f.MapFont()
	if f.OptionsData.ErrorMsg != "" {
		fmt.Println(f.OptionsData.ErrorMsg)
		return
	}
	// Split the input string and Build the output
	f.Split()

	f.OutputBuilder()
	if f.OptionsData.ErrorMsg != "" {
		fmt.Println(f.OptionsData.ErrorMsg)
		return
	}
	// Deliver the output to the console
	f.OutputDeliver()
	if f.OptionsData.ErrorMsg != "" {
		fmt.Println(f.OptionsData.ErrorMsg)
		return
	}
}
