package core

import (
	"fmt"
	"time"

	"github.com/iVitaliya/logger-go/utils"
)

func buildDate(t string, usesAMandPM bool) string {
	var result string

	switch t {
	case SHORT_TIMESTAMP_FORMAT:
		result = utils.ArrayReplace(
			SHORT_TIMESTAMP_FORMAT,
			[]string{
				DAY,
				MONTH,
				YEAR,
				TIME_SHORT,
			},
			[]string{
				fmt.Sprint(time.Now().Day()),
				"/" + fmt.Sprint(time.Now().Month()),
				"/" + fmt.Sprint(time.Now().Year()),
				" | " + fmt.Sprint(time.Now().Hour()) + ":" + fmt.Sprint(time.Now().Minute()),
			},
		)
		break
	case MID_TIMESTAMP_FORMAT:
		result = utils.ArrayReplace(
			MID_TIMESTAMP_FORMAT,
			[]string{
				DAY,
				MONTH,
				YEAR,
				TIME_LONG,
			},
			[]string{
				fmt.Sprint(time.Now().Day()),
				" " + fmt.Sprint(time.Now().Month()),
				" " + fmt.Sprint(time.Now().Year()),
				" | " + fmt.Sprint(time.Now().Hour()) + ":" + fmt.Sprint(time.Now().Minute()) + ":" + fmt.Sprint(time.Now().Second()),
			},
		)
		break
	case LONG_TIMESTAMP_FORMAT:
		if usesAMandPM {
			panic("AM/PM is not supported through the config, change either the timestamp format or enable AM/PM in the logger config")
		}
		result = utils.ArrayReplace(
			LONG_TIMESTAMP_FORMAT,
			[]string{
				DAY,
				MONTH,
				YEAR,
				TIME_AM_PM,
			},
			[]string{
				utils.DayFormatter(time.Now().Day()),
				" of " + utils.MonthFormatter(time.Now().Month()),
				utils.IntToString(time.Now().Year()),
				" | " + utils.AmOrPMFormatter(time.Now().Hour(), time.Now().Minute(), time.Now().Second()),
			},
		)
		break
	}

	return result
}

func prefix(config NamedLoggerConfiguration) string {
	
}

func processNamedConfig(config NamedLoggerConfiguration) {
	var (
		severity         = config.Severity
		name             = config.Name
		timestamp        = buildDate(config.TimestampFormat, config.TimeUsesAMandPM)
		outputFile       = config.OutputFile
		tags             = config.Tags
		enableColors     = config.EnableColors
		enableProcess    = config.EnableProccess
		displayProcessId = config.DisplayProccessId
	)
}
