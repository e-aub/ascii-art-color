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
	if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		return
	}
	f.ArgsChecker()
	if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		return
	}

	if f.Params.Input == "" {
		return
	}

	f.ToColorIndexes()

	// Minimize the input string
	f.Minimize()
	if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		return
	}
	// Map the input string to the selected font
	f.MapFont()
	if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		return
	}
	// Split the input string and Build the output
	f.Split()

	f.OutputBuilder()
	if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		return
	}
	// Deliver the output to the console
	f.OutputDeliver()
	if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		return
	}
}
