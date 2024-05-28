package functions

// Options struct contains the options for the program
type Parameters struct {
	Args           []string
	Input          string
	SplicedInput   []string
	ToMap          []rune
	OutputFile     string
	Output         string
	Color          string
	ToColor        string
	ToColorIndexes [][]int
	Banner         string
	ErrorMsg       string
}

var Params Parameters
