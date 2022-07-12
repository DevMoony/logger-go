package logger

import (
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
		Info:    utils.Info,
		Debug:   utils.Debug,
		Warning: utils.Warning,
		Error:   utils.Error,
		Fatal:   utils.Fatal,
		Log:     utils.Log(config.Severity),
	}

	return logs
}

func Info(message string) {
	utils.Info(message)
}

func Debug(message string) {
	utils.Debug(message)
}

func Warning(message string) {
	utils.Warning(message)
}

func Error(message string) {
	utils.Error(message)
}

func Fatal(message string, exitCode int) {
	utils.Fatal(message)
}
