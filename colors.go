// The colors package provide a simple way to bring colorful charcaters to terminal interface.
//
// This example will output the text with a Blue foreground and a Black background
//      colors.Println("@{bK}Example Text")
//
// This one will output the text with a red foreground
//      colors.Println("@rExample Text")
//
// This one will escape the @
//      colors.Println("@@")
//
// Full color syntax code
//      @{rgbcmykwRGBCMYKW}  foreground/background color
//      @{|}  Reset format style
//      @{!./_} Bold / Dim / Italic / underline
//      @{^&} Blink / Fast blink
//      @{?} Reverse the foreground and background color
//      @{-} Hide the text
// Note some of the functions are not widely supported, like "Fast blink" and "Italic".
package colors

import (
	"bytes"
	"fmt"
	"log"
)

// Escape character for color syntax
const escapeChar = '@'

// Short for reset to default style
var resetChar = fmt.Sprintf("%c|", escapeChar)

// Mapping from character to concrete escape code.
var codeMap = map[int]int{
	'|': 0,
	'!': 1,
	'.': 2,
	'/': 3,
	'_': 4,
	'^': 5,
	'&': 6,
	'?': 7,
	'-': 8,

	'k': 30,
	'r': 31,
	'g': 32,
	'y': 33,
	'b': 34,
	'm': 35,
	'c': 36,
	'w': 37,
	'd': 39,

	'K': 40,
	'R': 41,
	'G': 42,
	'Y': 43,
	'B': 44,
	'M': 45,
	'C': 46,
	'W': 47,
	'D': 49,
}

// Compile color syntax string like "rG" to escape code.
func colorMap(x string) string {
	attr := 0
	fg := 39
	bg := 49

	for _, key := range x {
		c, ok := codeMap[int(key)]
		switch {
		case !ok:
			log.Fatalf("Wrong color syntax: %c", key)
		case 0 <= c && c <= 8:
			attr = c
		case 30 <= c && c <= 37:
			fg = c
		case 40 <= c && c <= 47:
			bg = c
		}
	}
	return fmt.Sprintf("\033[%d;%d;%dm", attr, fg, bg)
}

// Handle state after meeting one '@'
func compileColorSyntax(input, output *bytes.Buffer) {
	i, _, err := input.ReadRune()
	if err != nil {
		// EOF got
		log.Fatal("Parse failed on color syntax")
	}

	switch i {
	default:
		output.WriteString(colorMap(string(i)))
	case '{':
		color := bytes.NewBufferString("")
		for {
			i, _, err := input.ReadRune()
			if err != nil {
				log.Fatal("Parse failed on color syntax")
			}
			if i == '}' {
				break
			}
			color.WriteRune(i)
		}
		output.WriteString(colorMap(color.String()))
	case escapeChar:
		output.WriteRune(escapeChar)
	}
}

// Compile the string and replace color syntax with concrete escape code.
func compile(x string) string {
	if x == "" {
		return ""
	}

	input := bytes.NewBufferString(x)
	output := bytes.NewBufferString("")

	for {
		i, _, err := input.ReadRune()
		if err != nil {
			break
		}
		switch i {
		default:
			output.WriteRune(i)
		case escapeChar:
			compileColorSyntax(input, output)
		}
	}
	return output.String()
}

// Compile multiple values, only do compiling on string type.
func compileValues(a *[]interface{}) {
	for i, x := range *a {
		if str, ok := x.(string); ok {
			(*a)[i] = compile(str)
		}
	}
}

// Similar to fmt.Print, will reset the color at the end.
func Print(a ...interface{}) (int, error) {
	a = append(a, resetChar)
	compileValues(&a)
	return fmt.Print(a...)
}

// Similar to fmt.Println, will reset the color at the end.
func Println(a ...interface{}) (int, error) {
	a = append(a, resetChar)
	compileValues(&a)
	return fmt.Println(a...)
}

// Similar to fmt.Printf, will reset the color at the end.
func Printf(format string, a ...interface{}) (int, error) {
	format += resetChar
	format = compile(format)
	return fmt.Printf(format, a...)
}

// Similar to fmt.Sprint, will reset the color at the end.
func Sprint(a ...interface{}) string {
	a = append(a, resetChar)
	compileValues(&a)
	return fmt.Sprint(a...)
}

// Similar to fmt.Sprintf, will reset the color at the end.
func Sprintf(format string, a ...interface{}) string {
	format += resetChar
	format = compile(format)
	return fmt.Sprintf(format, a...)
}
