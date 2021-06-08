package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	Println      = printlnInColor(nil)
	PrintlnLog   = printlnInColor(color.New(color.FgBlack, color.Faint).PrintfFunc())
	PrintlnInfo  = printlnInColor(color.White)
	PrintlnWarn  = printlnInColor(color.Yellow)
	PrintlnError = printlnInColor(color.Red)

	FPrintln      = fPrintlnInColor(nil)
	FPrintlnLog   = fPrintlnInColor(color.New(color.FgBlack, color.Faint).PrintfFunc())
	FPrintlnInfo  = fPrintlnInColor(color.White)
	FPrintlnWarn  = fPrintlnInColor(color.Yellow)
	FPrintlnError = fPrintlnInColor(color.Red)

	FPrint      = fPrintInColor(nil)
	FPrintLog   = fPrintInColor(color.New(color.FgBlack, color.Faint).PrintfFunc())
	FPrintInfo  = fPrintInColor(color.White)
	FPrintWarn  = fPrintInColor(color.Yellow)
	FPrintError = fPrintInColor(color.Red)
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

func fPrintInColor(color func(string, ...interface{})) func(string, ...interface{}) {
	printer := func(format string, args ...interface{}) {
		if color != nil {
			color(fmt.Sprintf(format, args...))
		} else {
			fmt.Printf(format, args...)
		}
	}
	return printer
}
