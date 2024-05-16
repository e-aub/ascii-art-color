package functions

func OutputBuilder(toPrint []string) {
	result := ""
	for _, part := range toPrint {
		if part == "\\n" {
			result += "\\n"
			continue
		}
		count := 0
		for count < 8 {
			for _, letter := range part {
				result += Font[letter][count]
			}
			result += "\\n"
			count++
		}
	}
	OptionsData.Output = result
}
