package functions

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RandomColor() {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	Colors["random"] = fmt.Sprintf("%d;%d;%d", generator.Intn(255), generator.Intn(255), generator.Intn(255))
}

func ToColorIndexes() {

	if Params.ToColor == "" {
		Params.ToColorIndexes = nil
		return
	}

	for i := 0; i < len(Params.Input)-len(Params.ToColor)+1; i++ {
		if Params.Input[i:i+len(Params.ToColor)] == Params.ToColor {
			Params.ToColorIndexes = append(Params.ToColorIndexes, []int{i, i + len(Params.ToColor) - 1})
			i += len(Params.ToColor) - 1
		}
	}
}

func InRange(index int) bool {
	for _, pair := range Params.ToColorIndexes {
		if index >= pair[0] && index <= pair[1] {
			return true
		}
	}
	return false
}

func HexToRgb(hexColor string) {
	if HexCheck.MatchString(hexColor) {
		r, err := strconv.ParseInt(hexColor[1:3], 16, 64)
		if err != nil {
			Params.Err = err
			return
		}
		g, err := strconv.ParseInt(hexColor[3:5], 16, 64)
		if err != nil {
			Params.Err = err
			return
		}
		b, err := strconv.ParseInt(hexColor[5:7], 16, 64)
		if err != nil {
			Params.Err = err
			return
		}
		Params.Color = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		return
	}
	Params.Err = errors.New("invalidhex")
}

func RGB(color string) {
	if match := RgbCheck.FindStringSubmatch(color); match != nil {
		r, err := strconv.Atoi(match[1])
		if err != nil {
			Params.Err = err
			return
		}
		g, err := strconv.Atoi(match[2])
		if err != nil {
			Params.Err = err
			return
		}
		b, err := strconv.Atoi(match[3])
		if err != nil {
			Params.Err = err
			return
		}
		if r > 255 || g > 255 || b > 255 {
			Params.Err = errors.New("rgbValue")
			return
		}
		Params.Color = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		return
	}
	Params.Err = errors.New("rgbFormat")
}
