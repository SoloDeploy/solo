package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	PrintlnLog    = printlnInColor(color.New(color.FgBlack, color.Faint).PrintfFunc())
	PrintlnInfo   = printlnInColor(nil)
	PrintlnWarn   = printlnInColor(color.Yellow)
	PrintlnError  = printlnInColor(color.Red)
	FPrintlnLog   = fPrintlnInColor(color.New(color.FgBlack, color.Faint).PrintfFunc())
	FPrintlnInfo  = fPrintlnInColor(nil)
	FPrintlnWarn  = fPrintlnInColor(color.Yellow)
	FPrintlnError = fPrintlnInColor(color.Red)
)

func printlnInColor(color func(string, ...interface{})) func(...interface{}) {
	printer := func(args ...interface{}) {
		if color != nil {
			color(fmt.Sprintln(args...))
		} else {
			fmt.Println(args...)
		}
	}
	return printer
}

func fPrintlnInColor(color func(string, ...interface{})) func(string, ...interface{}) {
	printer := func(format string, args ...interface{}) {
		if color != nil {
			color(fmt.Sprintln(fmt.Sprintf(format, args...)))
		} else {
			fmt.Println(fmt.Sprintf(format, args...))
		}
	}
	return printer
}
