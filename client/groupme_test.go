package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/client"
	"os"
	"time"
)

var _ = Describe("Groupme", func() {

	var (
		validClient    *groupme.BotClient
		testDay        time.Time
		testDayEnd     time.Time
		dayWithNoMemes time.Time
	)

	BeforeEach(func() {
		validClient = groupme.NewBotClient(os.Getenv("BOT_ID"), os.Getenv("GROUP_ID"), os.Getenv("GROUPME_ACCESS_TOKEN"))
		testDay = time.Date(2023, 11, 5, 0, 0, 0, 0, time.Local)
		testDayEnd = time.Date(testDay.Year(), testDay.Month(), 6, 0, 0, 0, 0, time.Local)
		dayWithNoMemes = time.Date(2023, 6, 25, 0, 0, 0, 0, time.UTC)
	})

	It("can throw an error if given incorrect credentials", func() {
		invalidClient := groupme.NewBotClient("fake-bot-id", "fake-groupid", "fake-accesstoken")

		messages, err := invalidClient.GetMemesInWindow(&testDay, &testDayEnd)
		Expect(err).To(HaveOccurred())
		Expect(messages).To(BeNil())
	})

	It("can query the groupme API to get messages", func() {
		messages, err := validClient.GetMemesInWindow(&testDay, &testDayEnd)

		Expect(err).NotTo(HaveOccurred())
		Expect(len(messages)).To(BeNumerically(">", 0))
		ExpectAllMessagesToBeOnDay(messages, testDay)
	})

	It("can upload an image to the image service", func() {
		imageReader, err := os.Open("client/testfiles/test_image.png")
		Expect(err).NotTo(HaveOccurred())
		defer imageReader.Close()

		res, err := validClient.ProcessImage(imageReader)
		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
		Expect(res).NotTo(BeEmpty())
	})

	It("can get the number of favorites for a meme", func() {
		message := groupme.Message{
			FavoritedBy: []string{
				"123",
				"456",
				"789",
			},
		}
		Expect(message.NumFavorites()).To(BeEquivalentTo(3))
	})

	It("can get the Memes from the groupme api and return the top meme", func() {
		meme, err := validClient.GetTopMemeBetweenDates(&testDay, &testDayEnd)

		expectedMessageID := "169923063376006591"
		Expect(err).NotTo(HaveOccurred())
		Expect(meme).NotTo(BeNil())
		Expect(meme.ID).To(BeEquivalentTo(expectedMessageID))
		ExpectAllMessagesToBeOnDay([]groupme.Message{meme}, testDay)
	})

	It("can get safely return no memes if none were sent that day", func() {
		meme, err := validClient.GetTopMemeBetweenDates(&dayWithNoMemes, &dayWithNoMemes)

		Expect(err).NotTo(HaveOccurred())
		Expect(meme).NotTo(BeNil())
		Expect(meme.ID).To(BeEmpty())
	})

})

func ExpectAllMessagesToBeOnDay(messages []groupme.Message, day time.Time) {
	for _, m := range messages {
		Expect(m.CreatedAt).To(And(BeNumerically(">=", day.Unix()), BeNumerically("<", day.AddDate(0, 0, 1).Unix())))
	}
}
