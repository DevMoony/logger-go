package utils

import (
	"fmt"
	"os"

	"github.com/iVitaliya/colors-go"
)

var log_time string = colorifyTime(colors.Yellow(Date()))
var print = func(msg ...any) { fmt.Println(msg) }

func Info(message ...any) {
	var state string = colorifyState("info")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(fmt.Sprintln(message)),
	}))
}

func Debug(message ...any) {
	var state string = colorifyState("debug")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(fmt.Sprintln(message)),
	}))
}

func Warning(message ...any) {
	var state string = colorifyState("warning")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(fmt.Sprintln(message)),
	}))
}

func Error(message ...any) {
	var state string = colorifyState("error")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(fmt.Sprintln(message)),
	}))
}

func Fatal(message ...any) {
	var state string = colorifyState("fatal")

	print(FormatString("{0} {1} {2}", []string{
		log_time,
		state,
		colors.BrightWhite(fmt.Sprintln(message)),
	}))

	os.Exit(1)
}

func Log(state string) func(...any) {
	var stater string = colorifyState(state)

	Log := func(message ...any) {
		print(FormatString("{0} {1} {2}", []string{
			log_time,
			stater,
			colors.BrightWhite(fmt.Sprint(message)),
		}))

		if state == "Fatal" {
			os.Exit(1)
		}
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
