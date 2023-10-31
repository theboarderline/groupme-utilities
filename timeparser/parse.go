package timeparser

import (
	"errors"
	"github.com/araddon/dateparse"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func FormatReportRange(s, e string) (*time.Time, *time.Time, error) {
	start, err := dateparse.ParseAny(s)
	if err != nil {
		return nil, nil, err
	}

	end, err := dateparse.ParseAny(e)
	if err != nil {
		return nil, nil, err
	}

	if start.Year() == 0 {
		now := time.Now()
		start = time.Date(now.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)
	}

	if end.Year() == 0 {
		now := time.Now()
		end = time.Date(now.Year(), end.Month(), end.Day(), 23, 59, 59, 0, time.UTC)
	}

	log.Debug().Msgf("start: %s, end: %s", start, end)
	return &start, &end, nil
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

func GetStartAndEndDateFromMessage(message string) (*time.Time, *time.Time, error) {

	splitMessage := strings.Split(message, " ")
	if len(splitMessage) == 2 {
		return nil, nil, errors.New("invalid date range")
	}

	return FormatReportRange(GetStartAndEndDateString(message))
}
