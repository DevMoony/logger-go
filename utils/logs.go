package utils

import (
	"fmt"
	"os"

	"github.com/iVitaliya/colors-go"
)

var log_time string = colorifyTime(colors.Yellow(Date()))
var print = func(msg string) { fmt.Println(msg) }

func Info(message string) {
	var state string = colorifyState("info")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(message),
	}))
}

func Debug(message string) {
	var state string = colorifyState("debug")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(message),
	}))
}

func Warning(message string) {
	var state string = colorifyState("warning")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(message),
	}))
}

func Error(message string) {
	var state string = colorifyState("error")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(message),
	}))
}

func Fatal(message string, exitCode int) {
	var state string = colorifyState("fatal")

	FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(message),
	})

	os.Exit(exitCode)
}

func Log(state string) func(string) {
	var stater string = colorifyState(state)

	Log := func(message string) {
		FormatString("{0} {1} {2}", []string{
			log_time,
			stater,
			colors.BrightWhite(message),
		})
	}

	return Log
}

func panic(reason string) {
	fmt.Println(stringAppender("", []string{
		log_time,
		" ",
		colors.Gray("("),
		colors.Red("PANIC"),
		colors.Gray(")"),
		" ",
		colors.BrightWhite(reason),
	}))

	os.Exit(1)
}
