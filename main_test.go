package main

import (
	f "ascii-art-fs/functions"
	"fmt"
	"log"
	"os"
	"testing"
)

func outputBuilder(toPrint []string) string {
	var result string
	for _, part := range toPrint {
		if part == "\\n" {
			fmt.Println()
			continue
		}
		count := 0
		for count < 8 {
			for _, letter := range part {
				result += (f.Font[letter][count])
			}
			result += "\n"
			count++
		}
	}
	return result
}

func mainCopy(toWrite string, banner string) string {
	if toWrite == "" {
		return ""
	}
	f.MapFont(banner, f.Minimize(toWrite))
	slicedToWrite := f.Split(toWrite)
	result := outputBuilder(slicedToWrite)
	return result
}

func TestMain(t *testing.T) {
	tests := []string{"0123456789", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz", "/(\")", "hello1\\nworld"}
	banners := []string{"standard", "shadow", "thinkertoy", "enigma"}
	fileName := "test_files/want%d%s.txt"
	var want string
	for i := 0; i < len(tests); i++ {
		for _, banner := range banners {
			got := mainCopy(tests[i], banner+".txt")
			want = readFile(fmt.Sprintf(fileName, i+1, banner))
			if string(got) != want {
				t.Errorf("Test case %d failed. Expected: %q", i+1, fmt.Sprintf(fileName, i+1, banner))
			}
		}
	}
}

func readFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
