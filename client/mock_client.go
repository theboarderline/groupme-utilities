package groupme

import (
	"io"
	"time"
)

type MockClient struct{}

var (
	fakeImage    = "https://i.groupme.com/1032x424.png.e509cc50c252443bb301dbba73a79701"
	fakeMessages = []Message{
		{
			ID:          "1",
			SenderID:    "1",
			FavoritedBy: []string{"1", "2", "3", "4"},
			Attachments: []Attachment{{
				Type: "image",
				URL:  fakeImage,
			}},
			CreatedAt: time.Date(2022, 6, 26, 12, 0, 0, 0, time.UTC).Unix(),
		},
		{
			ID:          "1",
			SenderID:    "1",
			FavoritedBy: []string{"1", "2", "3", "4"},
			Attachments: []Attachment{{
				Type: "image",
				URL:  fakeImage,
			}},
			CreatedAt: time.Date(2023, 6, 26, 12, 0, 0, 0, time.UTC).Unix(),
		},
		{
			ID:          "2",
			SenderID:    "1",
			FavoritedBy: []string{"1", "2", "3"},
			Attachments: []Attachment{{
				Type: "image",
				URL:  fakeImage,
			}},
			CreatedAt: time.Date(2023, 6, 26, 12, 0, 0, 0, time.UTC).Unix(),
		},
	}
)

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c MockClient) GetTopMemeBetweenDates(start, end *time.Time) (message Message, err error) {

	if start == nil || end == nil {
		return Message{}, nil
	}

	return fakeMessages[0], nil
}

func (c MockClient) GetMemesInWindow(start, end *time.Time) (messages []Message, err error) {

	return fakeMessages, nil
}

func (c MockClient) SendMessage(text, pictureURL string) error {
	return nil
}

func (c MockClient) ProcessImage(file io.Reader) (string, error) {
	return fakeImage, nil
}

func (c MockClient) GetReportForDateRange(start, end *time.Time) (report *Report, err error) {
	return NewReport(fakeMessages), nil
}
