package groupme

import (
	"io"
	"time"
)

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c MockClient) GetTopMemeBetweenDates(start, end *time.Time) (message Message, err error) {
	return Message{
		ID:          "1",
		SenderID:    "1",
		FavoritedBy: []string{"1", "2", "3", "4"},
		Attachments: []Attachment{{Type: "image"}},
	}, nil
}

func (c MockClient) GetMemesInWindow(start, end *time.Time) (messages []Message, err error) {

	return messages, nil
}

func (c MockClient) SendMessage(text, pictureURL string) error {
	return nil
}

func (c MockClient) ProcessImage(file io.Reader) (string, error) {
	return "https://fake-image.com", nil
}

func (c MockClient) GetReportForDateRange(start, end *time.Time) (report *Report, err error) {
	return &Report{
		StartDate: start,
		EndDate:   end,
	}, nil
}
