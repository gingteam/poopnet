package main

import (
	"math/rand"
	"strconv"
)

func RandomIntn(start int, end int) int {
	return rand.Intn(end-start) + start
}

func GenerateCaptcha(min int, max int) (int, int) {
	var numberOne int = RandomIntn(min, max)
	var numberTwo int = RandomIntn(min, max)
	return numberOne, numberTwo
}

func ColoredText(colorCode int, text string) string {
	foreground := "\033[" + strconv.Itoa(colorCode) + "m"
	resetColor := "\x1b[0m"
	return foreground + text + resetColor
}
