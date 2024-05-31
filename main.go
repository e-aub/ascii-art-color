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
	errHandler()
	f.ArgsChecker()
	errHandler()
	if f.Params.Input == "" {
		return
	}
	f.ToColorIndexes()
	// Minimize the input string
	f.Minimize()
	errHandler()
	// Map the input string to the selected font
	f.MapFont()
	errHandler()
	// Split the input string and Build the output
	f.Split()

	f.OutputBuilder()
	errHandler()
	// Deliver the output to the console
	f.OutputDeliver()
	errHandler()
}
func errHandler(){
if f.Params.Err != nil {
		fmt.Println(f.Errors[f.Params.Err.Error()])
		os.Exit(1)
	}
}
