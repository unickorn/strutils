package strutils

import (
	"fmt"
	"strings"
)

// cleaner is used to clean up minecraft color codes.
var cleaner *strings.Replacer

func init() {
	var minecraftToEmpty []string
	for i := 0; i < 10; i++ {
		minecraftToEmpty = append(minecraftToEmpty, fmt.Sprintf("§%d", i), "")
	}
	for _, v := range []string{"a", "b", "c", "d", "e", "f", "g", "k", "l", "o", "r"} {
		minecraftToEmpty = append(minecraftToEmpty, "§"+v, "")
	}
	cleaner = strings.NewReplacer(minecraftToEmpty...)
}

// Clean cleans a string from minecraft color codes.
func Clean(s string) string {
	return cleaner.Replace(s)
}
