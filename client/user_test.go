package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/client"
)

var _ = Describe("User", func() {

	It("can add a meme to a user", func() {
		user := groupme.User{
			ID:    "1",
			Name:  "test-user",
			Memes: groupme.NewMessageHeap(0),
		}
		user.AddMeme(groupme.Message{
			ID:          "1",
			FavoritedBy: []string{"1"},
			Attachments: []groupme.Attachment{
				{
					Type: "image",
				},
			},
		})
		Expect(user.NumMemes()).To(Equal(1))
		Expect(user.NumFavorites()).To(Equal(1))
	})

	It("can create a user from an existing message", func() {
		message := groupme.Message{
			ID:          "1",
			FavoritedBy: []string{"1"},
			Attachments: []groupme.Attachment{
				{
					Type: "image",
				},
			},
		}
		user := groupme.NewUser(message)
		Expect(user.NumMemes()).To(Equal(1))
		Expect(user.NumFavorites()).To(Equal(1))
	})
})
