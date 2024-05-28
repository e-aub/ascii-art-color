package main

import (
	f "ascii-art-output/functions"
	"fmt"
	"os"
)

func main() {
	// Check if the provided flags and args are valid
	f.Params.Args = os.Args[1:]
	f.FlagChecker()
	if f.Params.ErrorMsg != "" {
		fmt.Println(f.Params.ErrorMsg)
		return
	}
	f.ArgsChecker()
	if f.Params.ErrorMsg != "" {
		fmt.Println(f.Params.ErrorMsg)
		return
	}

	if f.Params.Input == "" {
		return
	}

	f.ToColorIndexes()

	// Minimize the input string
	f.Minimize()
	if f.Params.ErrorMsg != "" {
		fmt.Println(f.Params.ErrorMsg)
		return
	}
	// Map the input string to the selected font
	f.MapFont()
	if f.Params.ErrorMsg != "" {
		fmt.Println(f.Params.ErrorMsg)
		return
	}
	// Split the input string and Build the output
	f.Split()

	f.OutputBuilder()
	if f.Params.ErrorMsg != "" {
		fmt.Println(f.Params.ErrorMsg)
		return
	}
	// Deliver the output to the console
	f.OutputDeliver()
}
