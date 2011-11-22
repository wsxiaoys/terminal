package colors

import (
	"fmt"
)

const (
	// Attributes
	ATTR_RESET     = 0
	ATTR_BRIGHT    = 1
	ATTR_DIM       = 2
	ATTR_UNDERLINE = 3
	ATTR_BLINK     = 5
	ATTR_REVERSE   = 7

	// Foreground
	FG_BLACK   = 30
	FG_RED     = 31
	FG_GREEN   = 32
	FG_YELLOW  = 33
	FG_BLUE    = 34
	FG_MAGENTA = 35
	FG_CYAN    = 36
	FG_WHITE   = 37
	FG_DEFAULT = 39

	// Background
	BG_BLACK   = 40
	BG_RED     = 41
	BG_GREEN   = 42
	BG_YELLOW  = 43
	BG_BLUE    = 44
	BG_MAGENTA = 45
	BG_CYAN    = 46
	BG_WHITE   = 47
	BG_DEFAULT = 49
)

var (
	Default = Color{ATTR_RESET, FG_DEFAULT, BG_DEFAULT}
	Black   = Color{ATTR_RESET, FG_BLACK, BG_DEFAULT}
	Red     = Color{ATTR_RESET, FG_RED, BG_DEFAULT}
	Green   = Color{ATTR_RESET, FG_GREEN, BG_DEFAULT}
	Yellow  = Color{ATTR_RESET, FG_YELLOW, BG_DEFAULT}
	Blue    = Color{ATTR_RESET, FG_BLUE, BG_DEFAULT}
	Magenta = Color{ATTR_RESET, FG_MAGENTA, BG_DEFAULT}
	Cyan    = Color{ATTR_RESET, FG_CYAN, BG_DEFAULT}
	White   = Color{ATTR_RESET, FG_WHITE, BG_DEFAULT}
)

type Color struct {
	attr, fg, bg uint8
}

func (x Color) Format(f fmt.State, c rune) {
	if c == 'C' || c == 'v' {
		str := fmt.Sprint("\033[", x.attr, ";", x.fg, ";", x.bg, "m")
		bytes := []byte(str)
		f.Write(bytes)
	} else {
		f.Write([]byte(fmt.Sprint("#!C(colors.Color)")))
	}
}

func (x Color) Print(a ...interface{}) (n int, err error) {
	fmt.Print(x)
	n, err = fmt.Print(a...)
	fmt.Print(Default)
	return
}

func (x Color) Println(a ...interface{}) (n int, err error) {
	fmt.Print(x)
	n, err = fmt.Println(a...)
	fmt.Print(Default)
	return
}

func (x Color) Printf(format string, a ...interface{}) (n int, err error) {
	fmt.Print(x)
	n, err = fmt.Printf(format, a...)
	fmt.Print(Default)
	return
}

func Cprint(a ...interface{}) (n int, err error) {
	n, err = fmt.Print(a...)
	fmt.Print(Default)
	return
}

func Cprintf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Printf(format, a...)
	fmt.Print(Default)
	return
}
