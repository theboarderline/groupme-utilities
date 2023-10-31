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

	if start == nil || end == nil {
		return Message{}, nil
	}

	return Message{
		ID:          "1",
		SenderID:    "1",
		FavoritedBy: []string{"1", "2", "3", "4"},
		Attachments: []Attachment{{Type: "image"}},
		CreatedAt:   time.Date(2023, 6, 26, 12, 0, 0, 0, time.UTC).Unix(),
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
	messages := []Message{
		{
			ID:          "1",
			SenderID:    "1",
			FavoritedBy: []string{"1", "2", "3", "4"},
			Attachments: []Attachment{{Type: "image"}},
			CreatedAt:   time.Date(2023, 6, 26, 12, 0, 0, 0, time.UTC).Unix(),
		},
		{
			ID:          "2",
			SenderID:    "1",
			FavoritedBy: []string{"1", "2", "3"},
			Attachments: []Attachment{{Type: "image"}},
			CreatedAt:   time.Date(2023, 6, 26, 12, 0, 0, 0, time.UTC).Unix(),
		},
	}

	return NewReport(messages), nil
}
