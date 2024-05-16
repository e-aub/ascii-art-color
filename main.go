package main

import (
	f "ascii-art-fs/functions"
	"os"
)

func main() {
	f.ArgsChecker(os.Args)
	// if f.OptionsData.Input == "" {
	// 	return
	// }
	// f.MapFont(f.OptionsData.Banner, f.Minimize(toWrite))
	// slicedToWrite := f.Split(toWrite)
	// f.OutputBuilder(slicedToWrite)

}
