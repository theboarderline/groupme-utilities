package groupme

import (
	"time"
)

type Message struct {
	ID          string       `json:"id,omitempty"`
	SenderID    string       `json:"sender_id,omitempty"`
	Name        string       `json:"name"`
	Text        string       `json:"text"`
	FavoritedBy []string     `json:"favorited_by"`
	Attachments []Attachment `json:"attachments"`
	CreatedAt   int64        `json:"created_at,omitempty"`
}

type Attachment struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (m Message) SentDuringTimespan(begin, end time.Time) bool {
	return m.SentAfter(begin) && m.SentBefore(end)
}

func (m Message) SentBefore(day time.Time) bool {
	isBefore := m.CreatedAt < day.Unix()
	return isBefore
}

func (m Message) SentAfter(day time.Time) bool {
	isAfter := m.CreatedAt >= day.Unix()
	return isAfter
}

func (m Message) NumFavorites() int {
	return len(m.FavoritedBy)
}

func (m Message) IsMeme() bool {

	for _, attachment := range m.Attachments {
		if attachment.Type == "image" {
			return true
		}
	}

	return false
}
