package groupme_test

import (
	"github.com/theboarderline/groupme-utilities/client"
	"os"
	"time"
)

var _ = Describe("Groupme", func() {

	var (
		validClient    *groupme.Client
		testDay        time.Time
		testDayEnd     time.Time
		dayWithNoMemes time.Time
	)

	BeforeEach(func() {
		validClient = groupme.NewClient(os.Getenv("BOT_ID"), os.Getenv("GROUP_ID"), os.Getenv("GROUPME_ACCESS_TOKEN"))
		testDay = time.Date(2023, 6, 14, 0, 0, 0, 0, time.UTC)
		testDayEnd = time.Date(testDay.Year(), testDay.Month(), testDay.Day(), 23, 59, 59, 0, time.UTC)
		dayWithNoMemes = time.Date(2023, 6, 25, 0, 0, 0, 0, time.UTC)
	})

	It("can throw an error if given incorrect credentials", func() {
		invalidClient := groupme.NewClient("fake-bot-id", "fake-groupid", "fake-accesstoken")

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
		imageReader, err := os.Open("testfiles/test_image.png")
		Expect(err).NotTo(HaveOccurred())
		defer imageReader.Close()

		res, err := validClient.UploadPicture(imageReader)
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
		meme, err := validClient.GetTopMemeBetweenDates(testDay)

		expectedMessageID := "168675296063513419"
		Expect(err).NotTo(HaveOccurred())
		Expect(meme).NotTo(BeNil())
		Expect(meme.ID).To(BeEquivalentTo(expectedMessageID))
		ExpectAllMessagesToBeOnDay([]groupme.Message{meme}, testDay)
	})

	It("can get the Memes from the groupme api and return the top meme", func() {
		meme, err := validClient.GetTopMemeBetweenDates(testDay)

		expectedMessageID := "168675296063513419"
		Expect(err).NotTo(HaveOccurred())
		Expect(meme).NotTo(BeNil())
		Expect(meme.ID).To(BeEquivalentTo(expectedMessageID))
		ExpectAllMessagesToBeOnDay([]groupme.Message{meme}, testDay)
	})

	It("can get safely return no memes if none were sent that day", func() {
		meme, err := validClient.GetTopMemeBetweenDates(dayWithNoMemes)

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
