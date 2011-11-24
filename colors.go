package colors

import (
	"bytes"
	"fmt"
	"log"
)

const escapeCode = '@'

const resetCode = "@0"

var codeMap = map[int]int{
	'0': 0,
	'!': 1,
	'.': 2,
	'_': 3,
	'?': 5,
	'/': 7,
	'-': 8,

	'r': 31,
	'g': 32,
	'y': 33,
	'b': 34,
	'm': 35,
	'c': 36,
	'w': 37,
	'd': 39,

	'R': 41,
	'G': 42,
	'Y': 43,
	'B': 44,
	'M': 45,
	'C': 46,
	'W': 47,
	'D': 49,
}

func colorMap(x string) string {
	attr := 0
	fg := 39
	bg := 49

	for _, key := range x {
		c, ok := codeMap[key]
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
	case escapeCode:
		output.WriteRune(escapeCode)
	}
}

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
		case escapeCode:
			compileColorSyntax(input, output)
		}
	}
	return output.String()
}

func compileValues(a *[]interface{}) {
	for i, x := range *a {
                if str, ok := x.(string); ok {
			(*a)[i] = compile(str)
                }
	}
}

func Print(a ...interface{}) (int, error) {
	a = append(a, resetCode)
	compileValues(&a)
	return fmt.Print(a...)
}

func Println(a ...interface{}) (int, error) {
	a = append(a, resetCode)
	compileValues(&a)
	return fmt.Println(a...)
}

func Printf(format string, a ...interface{}) (int, error) {
	format += resetCode
	format = compile(format)
	return fmt.Printf(format, a...)
}
