package calculator

import (
	"strconv"
	"strings"
)

func isEmptyString(text string) bool {
	return text == ""
}

func isDoubleComma(text string) bool {
	return strings.Contains(text, ",,")
}

func Add(numbers string) int {
	sum := 0
	// check empty
	if isEmptyString(numbers) {
		return 0
	}

	// replace new line
	replacedNewline := strings.Replace(numbers, "\n", ",", -1)

	// check double comma
	if isDoubleComma(replacedNewline) {
		panic("Not need to prove it, just clarifying")
	}

	// slit string comma delimited
	numberslice := strings.Split(replacedNewline, ",")

	// sum all numbers
	for _, number := range numberslice {
		n, _ := strconv.Atoi(number)
		sum += n
	}
	return sum
}
