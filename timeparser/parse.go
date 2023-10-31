package timeparser

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func formatTime(s string) *time.Time {
	if s == "" {
		return nil
	}

	date, err := time.Parse("01/02", s)
	if err != nil {
		log.Debug().Msgf("Error parsing start date: %s", err)
	}

	currentYear := fmt.Sprint(time.Now().Year())
	date, err = time.Parse("01/02/2006", fmt.Sprintf("%s/%s", s, currentYear))

	return &date
}

func FormatReportRange(s, e string) (start *time.Time, end *time.Time) {
	start = formatTime(s)
	end = formatTime(e)

	return start, end
}

func GetStartAndEndDateString(date string) (start string, end string) {
	dateList := strings.Split(date, "-")
	if len(dateList) == 0 {
		return "", ""
	}
	if len(dateList) == 1 {
		return dateList[0], ""
	}

	return strings.TrimSpace(dateList[0]), strings.TrimSpace(dateList[1])
}

func GetStartAndEndDateFromMessage(message string) (start *time.Time, end *time.Time) {
	splitMessage := strings.Split(message, " ")
	if len(splitMessage) == 2 {
		return nil, nil
	}
	startString, endString := GetStartAndEndDateString(message)
	return FormatReportRange(startString, endString)
}
