package utils

import (
	"time"
)

// =================================
// ========== [[ CONV ]] ===========
// =================================

// amOrPMFormatter formats and returns the time in a 12-hour clock format with AM/PM suffix.
// It takes hour, minute, and second as input parameters and returns the formatted time string.
func AmOrPMFormatter(hour int, minute int, second int) string {
	var (
		status string
		mins   string
		secs   string
	)

	if hour >= 12 {
		status = "PM"
	} else if hour == 0 || hour < 12 {
		status = "AM"
	}

	if minute < 10 {
		mins = stringAppender("0", []string{
			IntToString(minute),
		})
	} else if minute > 9 {
		mins = IntToString(minute)
	}

	if second < 10 {
		secs = stringAppender("0", []string{
			IntToString(second),
		})
	} else if second > 9 {
		secs = IntToString(second)
	}

	return FormatString("{0}:{1}:{2} {3}", []string{
		IntToString(hour),
		mins,
		secs,
		status,
	})
}

func TimeFormatter(hour int, minute int, include_seconds bool) string {
	var (
		mins string
		secs string
	)

	if include_seconds {
	}
}

// dayFormatter takes a day of the month as an integer and returns a string in the format [number] [suffix].
// The suffix is determined by the day of the month, following the standard English rules for ordinal numbers.
// For example, the 1st, 2nd, 3rd, 21st, 22nd, and 23rd will have the suffixes "st", "nd", "rd", "st", "nd", and "rd", respectively.
// All other days of the month will have the suffix "th".
func DayFormatter(day int) string {
	var result string

	if day == 1 || day == 21 || day == 31 {
		result = IntToString(day) + "st"
	} else if day == 2 || day == 22 {
		result = IntToString(day) + "nd"
	} else if day == 3 || day == 23 {
		result = IntToString(day) + "rd"
	} else {
		result = IntToString(day) + "th"
	}

	return result
}

// monthFormatter takes a time.Month and returns a string of the full name of the month, such as "January" or "February".
func MonthFormatter(month time.Month) string {
	var result string

	switch month {
	case time.January:
		result = "January"

	case time.February:
		result = "February"

	case time.March:
		result = "March"

	case time.April:
		result = "April"

	case time.May:
		result = "May"

	case time.June:
		result = "June"

	case time.July:
		result = "July"

	case time.August:
		result = "August"

	case time.September:
		result = "September"

	case time.October:
		result = "October"

	case time.November:
		result = "November"

	case time.December:
		result = "December"
	}

	return result
}

// =================================
// ========== [[ TIME ]] ===========
// =================================

func Date() string {
	var (
		str string = ""
		now        = time.Now()
	)
	var (
		amOrPm = AmOrPMFormatter(now.Hour(), now.Minute(), now.Second())
		day    = DayFormatter(now.Day())
		month  = MonthFormatter(now.Month())
		year   = IntToString(now.Year())
	)

	str = FormatString("{0} of {1} {2} | {3}", []string{
		day,
		month,
		year,
		amOrPm,
	})

	return str
}
