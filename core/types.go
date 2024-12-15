package core

import "io"

type NamedLoggerConfiguration struct {
	Severity          string   // Minimum Severity level to log.
	Name              string   // Name for the logger. (e.g., service name)
	TimestampFormat   string   // Timestamp format to use for logging. uses place-holders: <YEAR>, <MONTH>, <DAY>, <HOUR>, <MINUTE>, <SECOND>, <AM/PM>. (e.g.: <DAY> <MONTH> <YEAR> | <HOUR>:<MINUTE>:<SECOND> <AM/PM>)
	OutputFile        string   // Filepath for writing logs to. (e.g.: "logs") NOTE: <YEAR>, <MONTH>, <DAY>, <HOUR>, <MINUTE>, <SECOND>, <AM/PM> are usable placeholders and will be replaced with the current date.
	EnableColors      bool     // Whether or not colors should be enabled for logging.
	Tags              []string // Custom tags for the logger.
	EnableProccess    bool     // Whether or not the logger should log process information.
	DisplayProccessId bool     // Whether or not the logger should display the current process id in the logs.
	TimeUsesAMandPM   bool     // Whether or not the logger should display the current time in 12-hour format with AM/PM suffix or not.
}

type LoggerConfiguration struct {
	Severity          string   // Minimum Severity level to log.
	TimestampFormat   string   // Timestamp format to use for logging. uses place-holders: <YEAR>, <MONTH>, <DAY>, <HOUR>, <MINUTE>, <SECOND>, <AM/PM>. (e.g.: <DAY> <MONTH> <YEAR> | <HOUR>:<MINUTE>:<SECOND> <AM/PM>)
	OutputFile        string   // Filepath for writing logs to. (e.g.: "logs") NOTE: <YEAR>, <MONTH>, <DAY>, <HOUR>, <MINUTE>, <SECOND>, <AM/PM> are usable placeholders and will be replaced with the current date.
	EnableColors      bool     // Whether or not colors should be enabled for logging.
	Tags              []string // Custom tags for the logger.
	EnableProccess    bool     // Whether or not the logger should log process information.
	DisplayProccessId bool     // Whether or not the logger should display the current process id in the logs.
	TimeUsesAMandPM   bool     // Whether or not the logger should display the current time in 12-hour format with AM/PM suffix or not
}

type Logger interface {
	// Logs a message to the terminal with the given configuration severity (log) level.
	Log(message ...any)
	// Logs a info message to the terminal.
	Info(message ...any)
	// Logs a debug message to the terminal.
	Debug(message ...any)
	// Logs a warning message to the terminal.
	Warning(message ...any)
	// Logs a error message to the terminal.
	Error(message ...any)
	// Logs a fatal message to the terminal and exits the program.
	Fatal(message ...any)

	Scan(a ...any) (n int, err error)
	Scanf(format string, a ...any) (n int, err error)
	Scanln(a ...any) (n int, err error)

	Sscan(str string, a ...any) (n int, err error)
	Sscanf(str string, format string, a ...any) (n int, err error)
	Sscanln(str string, a ...any) (n int, err error)

	Fscan(r io.Reader, a ...any) (n int, err error)
	Fscanf(r io.Reader, format string, a ...any) (n int, err error)
	Fscanln(r io.Reader, a ...any) (n int, err error)
}
