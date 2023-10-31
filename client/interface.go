package groupme

import (
	"io"
	"time"
)

type Client interface {
	GetTopMemeBetweenDates(start, end *time.Time) (message Message, err error)
	GetMemesInWindow(start, end *time.Time) (messages []Message, err error)
	SendMessage(text, pictureURL string) error
	ProcessImage(file io.Reader) (string, error)
	GetReportForDateRange(start, end *time.Time) (report *Report, err error)
}
