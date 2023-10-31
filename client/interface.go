package groupme

import (
	"io"
	"time"
)

type Client interface {
	GetTopMemeBetweenDates(startDate, endDate time.Time) (message Message, err error)
	GetMemesInWindow(startDate, endDate *time.Time) (messages []Message, err error)
	SendMessage(text, pictureURL string) error
	ProcessImage(file io.Reader) (string, error)
}
