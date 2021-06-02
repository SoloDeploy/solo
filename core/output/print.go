package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	PrintInfo   = printInColor(color.Blue)
	PrintWarn   = printInColor(color.Yellow)
	PrintError  = printInColor(color.Red)
	PrintfInfo  = printfInColor(color.Blue)
	PrintfWarn  = printfInColor(color.Yellow)
	PrintfError = printfInColor(color.Red)
)

func printInColor(color func(string, ...interface{})) func(...interface{}) {
	printer := func(args ...interface{}) {
		color(fmt.Sprint(args...))
	}
	return printer
}

func printfInColor(color func(string, ...interface{})) func(string, ...interface{}) {
	printer := func(format string, args ...interface{}) {
		color(fmt.Sprintf(format, args...))
	}
	return printer
}
