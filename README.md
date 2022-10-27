# strutils
Center text or strip color codes for Minecraft.

## Usage
```go
text := "§r§fThis is a multiline text\n§aThat needs to be §lcentered§r\n§eSo that it looks §ocool§r§e in Minecraft"

// Remove color codes from a string
strutils.Clean(text)

// CenterLine(text)
// center a mutliline string to fit in a line of chat in Minecraft.
strutils.CenterLine(text)

// Center(text, maxLength, addRightPadding)
// center with auto-detection of maximum length and add no right padding to lines
strutils.Center(text, 0, false)

// optionally, set your own maximum pixel length to make the lines even longer
// a line in minecraft is around 180 pixels long, so this may not look very good in the game
strutils.Center(text, 200, false)

```
