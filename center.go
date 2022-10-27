package strutils

// https://github.com/NiclasOlofsson/MiNET/blob/master/src/MiNET/MiNET/Utils/TextUtils.cs

import (
	"math"
	"strings"
)

const (
	lineLength = 30
	charWidth  = 6
	spaceChar  = ' '
)

var charWidths = map[string]int{
	" ":  4,
	"!":  2,
	"\"": 5,
	"'":  3,
	"(":  5,
	")":  5,
	"*":  5,
	",":  2,
	".":  2,
	":":  2,
	";":  2,
	"<":  5,
	">":  5,
	"@":  7,
	"I":  4,
	"[":  4,
	"]":  4,
	"f":  5,
	"i":  2,
	"k":  5,
	"l":  3,
	"t":  4,
	"|":  2,
	"~":  7,
	"█":  9,
	"░":  8,
	"▒":  9,
	"▓":  9,
	"▌":  5,
	"─":  9,
}

// CenterLine centers a string on a regular line in minecraft.
func CenterLine(input string) string {
	return Center(input, lineLength*charWidth, false)
}

// Center centers a [multiline] string with a given maximum length. addRightPadding can be set to true
// to add padding to the right of the centered lines.
func Center(input string, maxLength int, addRightPadding bool) string {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l := getPixelLength(line)
		if l > maxLength {
			maxLength = l
		}
	}
	result := strings.Builder{}
	spaceWidth := 4
	for _, line := range lines {
		length := math.Max(float64(maxLength-getPixelLength(line)), 0)
		padding := int(math.Round(float64(int(length) / (2 * spaceWidth))))
		result.WriteString(strings.Repeat(string(spaceChar), padding) + line)
		if addRightPadding {
			paddingRight := int(math.Floor(float64(int(length) / (2 * spaceWidth))))
			result.WriteString(strings.Repeat(string(spaceChar), paddingRight))
		}
		result.WriteString("\n")
	}
	return strings.TrimRight(result.String(), "\n")
}

// getCharWidth returns the width of a character.
func getCharWidth(s string) int {
	c, ok := charWidths[s]
	if !ok {
		return charWidth
	}
	return c
}

// getPixelLength returns the total pixel length of a string.
func getPixelLength(line string) int {
	length := 0
	for _, c := range Clean(line) {
		length += getCharWidth(string(c))
	}

	// +1 for each bold character
	split := strings.Split(line, "§")
	bold := false
	for _, s := range split {
		if len(s) == 0 {
			continue
		}
		switch s[0] {
		case 'l':
			bold = true
		case 'r':
			bold = false
		}
		if bold {
			length += len(s) - 1
		}
	}
	return length
}
