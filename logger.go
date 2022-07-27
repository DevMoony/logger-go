package logger

import (
	"fmt"

	"github.com/iVitaliya/logger-go/utils"
)

const (
	InfoState    string = "Info"
	DebugState   string = "Debug"
	WarningState string = "Warning"
	ErrorState   string = "Error"
	FatalState   string = "Fatal"
)

type LoggerConfig struct {
	Severity string
}

type ILogger struct {
	Info    func(string)
	Debug   func(string)
	Warning func(string)
	Error   func(string)
	Fatal   func(string)
	Log     func(string)
}

func ConfigureLogger(config LoggerConfig) *ILogger {
	logs := &ILogger{
		// Logs a info message to the terminal.
		Info: utils.Info,
		// Logs a debug message to the terminal.
		Debug: utils.Debug,
		// Logs a warning message to the terminal.
		Warning: utils.Warning,
		// Logs a error message to the terminal.
		Error: utils.Error,
		// Logs a fatal message to the terminal and exits the program.
		Fatal: utils.Fatal,
		// Logs a message to the terminal with the given configuration.
		Log: utils.Log(config.Severity),
	}

	return logs
}

func LogEmptySpace(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print(" ")
	}
}

func LogNextLine(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print("\n")
	}
}

// Logs a info message to the terminal.
func Info(message string) {
	utils.Info(message)
}

// Logs a debug message to the terminal.
func Debug(message string) {
	utils.Debug(message)
}

// Logs a warning message to the terminal.
func Warning(message string) {
	utils.Warning(message)
}

// Logs a error message to the terminal.
func Error(message string) {
	utils.Error(message)
}

// Logs a fatal message to the terminal and exits the program.
func Fatal(message string, exitCode int) {
	utils.Fatal(message)
}

// Scanf scans text read from standard input, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why. Newlines in the input must match newlines in the format. The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline.
func Scanf(format string, a ...any) (n int, err error) {
	return fmt.Scanf(format, a)
}

// Scan scans text read from standard input, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.
func Scan(a ...any) (n int, err error) {
	return fmt.Scan(a)
}

// Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
func Scanln(a ...any) (n int, err error) {
	return fmt.Scanln(a)
}

// Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.
func Sscanf(str string, format string, a ...any) (n int, err error) {
	return fmt.Sscanf(str, format, a)
}

// Sscan scans the argument string, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.
func Sscan(str string, a ...any) (n int, err error) {
	return fmt.Sscan(str, a)
}

// Sscanln is similar to Sscan, but stops scanning at a newline and after the final item there must be a newline or EOF.
func Sscanln(str string, a ...any) (n int, err error) {
	return fmt.Sscanln(str, a)
}
