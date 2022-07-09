package utils

import "github.com/iVitaliya/colors-go"

func colorifyTime(value string) string {
	str := FormatString("{0}{1}{2}", []string{
		colors.Gray("["),
		value,
		colors.Gray("]"),
	})

	return str
}

func colorifyState(severity string) string {
	lower := ToLower(severity)
	handleLogSeverityError(lower)

	identifySeverity := func(sev string) string {
		var s string

		switch sev {
		case "info":
			s = colors.BrightBlue("Info")
		case "debug":
			s = colors.Green("Debug")
		case "warn":
		case "warning":
			s = colors.BrightYellow("Warning")
		case "error":
			s = colors.BrightRed("Error")
		case "fatal":
			s = colors.Red("Fatal")
		}

		return s
	}

	str := FormatString("{0}{1}{2}", []string{
		colors.Gray("("),
		identifySeverity(lower),
		colors.Gray(")"),
	})

	return str
}

func handleLogSeverityError(severity string) {
	availables := ArrayIncludes([]string{
		"info",
		"debug",
		"warn",
		"warning",
		"error",
		"fatal",
	}, severity)
	if !availables {
		panic(stringAppender("", []string{
			"The value \"",
			severity,
			"\" isn't a valid severity level for the logger.",
		}))
	}
}
