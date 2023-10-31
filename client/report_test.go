package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/client"
)

var _ = Describe("Report", func() {

	var (
		report *groupme.Report
	)

	BeforeEach(func() {
		memes := []groupme.Message{
			{
				ID:          "1",
				SenderID:    "1",
				FavoritedBy: []string{"1", "2", "3", "4"},
				Attachments: []groupme.Attachment{{Type: "image"}},
			},
			{
				ID:          "2",
				SenderID:    "1",
				FavoritedBy: []string{"1", "2", "3"},
				Attachments: []groupme.Attachment{{Type: "image"}},
			},
			{
				ID:          "3",
				SenderID:    "2",
				FavoritedBy: []string{"1", "2"},
				Attachments: []groupme.Attachment{{Type: "image"}},
			},
		}

		report = groupme.NewReport(memes)
		Expect(report.MemesHeap.Len()).To(Equal(3))
		Expect(report.UsersHeap.Len()).To(Equal(2))

	})

	It("can get a meme report for a set of memes", func() {
		actualUser := report.PopUser()
		Expect(actualUser.ID).To(Equal("1"))
		Expect(actualUser.NumFavorites()).To(Equal(7))

		orderedMemes := report.MemesHeap.List()
		Expect(orderedMemes).To(HaveLen(3))
		Expect(orderedMemes[0].ID).To(Equal("1"))
		Expect(orderedMemes[1].ID).To(Equal("2"))
		Expect(orderedMemes[2].ID).To(Equal("3"))
	})

	It("can format a message to send based on a report", func() {
		reportMessage := report.String()
		Expect(reportMessage).To(ContainSubstring("Meme Report:"))
	})

})
