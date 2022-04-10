package strcenter

import "strings"

var m = []string{
	"§0",
	"§1",
	"§2",
	"§3",
	"§4",
	"§5",
	"§6",
	"§7",
	"§8",
	"§9",
	"§a",
	"§b",
	"§c",
	"§d",
	"§e",
	"§f",
	"§g",
	"§k",
	"§l",
	"§o",
	"§r",
}

// cleaner is used to clean up minecraft color codes.
var cleaner *strings.Replacer

func init() {
	var minecraftToEmpty []string
	for _, c := range m {
		minecraftToEmpty = append(minecraftToEmpty, c, "")
	}
	cleaner = strings.NewReplacer(minecraftToEmpty...)
}

// Clean cleans a string from minecraft color codes.
func Clean(s string) string {
	return cleaner.Replace(s)
}
