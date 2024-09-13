package excelSheetColumnNumber

import (
	"math"
)

func TitleToNumber(columnTitle string) int {
	var str2Num [26]int
	for i := 0; i < 26; i++ {
		str2Num[i] = i + 1
	}

	result := 0
	l := len(columnTitle) - 1
	for idx, str := range columnTitle {
		if l > idx {
			result += str2Num[str-'A'] * 26 * int(math.Pow(26, float64(l - idx - 1)))
		} else {
			result += str2Num[str-'A']
		}
	}
	return result
}

func TitleToNumberOp(columnTitle string) int {
	multiplier := 1
	result := 0

	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += (int(columnTitle[i]) - int('A') + 1) * multiplier
		multiplier *= 26
	}

	return result
}
