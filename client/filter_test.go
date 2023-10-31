package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/client"
	"time"
)

var _ = Describe("Filter", func() {
	var (
		todayBegin time.Time
		todayEnd   time.Time
	)

	BeforeEach(func() {
		todayBegin = time.Date(2023, 6, 14, 0, 0, 0, 0, time.UTC)
		todayEnd = time.Date(todayBegin.Year(), todayBegin.Month(), todayBegin.Day(), 23, 59, 59, 0, time.UTC)
	})

	It("can determine if a message was sent in a given time window", func() {

	})

	It("can filter messages and return only the Memes on a given day", func() {
		yesterday := time.Date(2023, 6, 12, 0, 0, 0, 0, time.UTC)
		rawMessages := []groupme.Message{
			{
				Attachments: []groupme.Attachment{{Type: "image"}},
				CreatedAt:   todayBegin.Unix(),
			},
			{
				Attachments: []groupme.Attachment{{Type: "image"}},
				CreatedAt:   todayBegin.Unix(),
			},
			{
				Attachments: []groupme.Attachment{{Type: "not an image"}},
				CreatedAt:   todayBegin.Unix(),
			},
			{
				Attachments: []groupme.Attachment{{Type: "image"}},
				CreatedAt:   yesterday.Unix(),
			},
		}

		filteredMessages := groupme.FilterMemesByTimespan(rawMessages, &todayBegin, &todayEnd)

		Expect(len(filteredMessages)).To(BeEquivalentTo(2))
	})

	It("can determine if a message is a meme", func() {
		message := groupme.Message{
			Attachments: []groupme.Attachment{{Type: "image"}},
		}
		Expect(message.IsMeme()).To(BeTrue())
	})

	It("can determine if a message is not a meme", func() {
		message := groupme.Message{
			Attachments: []groupme.Attachment{{Type: "not an image"}},
		}
		Expect(message.IsMeme()).To(BeFalse())
	})

	It("can determine if a message was sent on a specific day", func() {
		message := groupme.Message{
			CreatedAt: todayBegin.Unix(),
		}
		sentDuringTimespan := groupme.MessageSentDuringTimespan(message, &todayBegin, &todayEnd)
		Expect(sentDuringTimespan).To(BeTrue())
	})

	It("can determine if a message was not sent on a specific day", func() {
		message := groupme.Message{
			CreatedAt: todayBegin.AddDate(0, 0, 1).Unix(),
		}
		sentDuringTimespan := groupme.MessageSentDuringTimespan(message, &todayBegin, &todayEnd)
		Expect(sentDuringTimespan).To(BeFalse())

		message = groupme.Message{
			CreatedAt: todayBegin.AddDate(0, 0, 1).Unix(),
		}
		sentDuringTimespan = groupme.MessageSentDuringTimespan(message, &todayBegin, &todayEnd)
		Expect(sentDuringTimespan).To(BeFalse())
	})

	It("can filter messages and return n Memes with the most likes", func() {
		rawMessages := []groupme.Message{
			{
				ID:          "1",
				Attachments: []groupme.Attachment{{Type: "image"}},
				FavoritedBy: []string{"1", "3"},
			},
			{
				ID:          "2",
				Attachments: []groupme.Attachment{{Type: "image"}},
				FavoritedBy: []string{"1", "2", "3"},
			},
			{
				ID:          "3",
				Attachments: []groupme.Attachment{{Type: "image"}},
				FavoritedBy: []string{"1"},
			},
		}

		memes := groupme.GetTopNMessages(rawMessages, 1)
		Expect(len(memes)).To(BeEquivalentTo(1))
		Expect(memes[0].ID).To(Equal("2"))
	})

})
