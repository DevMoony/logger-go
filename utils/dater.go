package utils

import (
	"time"
)

// =================================
// ========== [[ CONV ]] ===========
// =================================

func dayFormatter(day int) string {
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

func monthFormatter(month time.Month) string {
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
		day   = dayFormatter(now.Day())
		month = monthFormatter(now.Month())
		year  = IntToString(now.Year())
	)

	str = FormatString("{0} of {1} {2}", []string{
		day,
		month,
		year,
	})

	return str
}
