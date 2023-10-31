package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/client"
)

var _ = Describe("UserHeap", func() {

	It("can sort two users and return the one with most total favorites", func() {
		heap := groupme.NewUserHeap(0)
		Expect(heap).NotTo(BeNil())

		losingUser := groupme.User{
			ID:    "losing-user",
			Memes: groupme.NewMessageHeap(0),
		}
		losingUser.AddMeme(groupme.Message{
			ID:          "1",
			FavoritedBy: []string{"1"},
			Attachments: []groupme.Attachment{{Type: "image"}},
		})
		Expect(losingUser.NumFavorites()).To(Equal(1))

		middleUser := groupme.User{
			ID:    "middle-user",
			Memes: groupme.NewMessageHeap(0),
		}
		middleUser.AddMeme(groupme.Message{
			ID:          "2",
			FavoritedBy: []string{"1", "2"},
			Attachments: []groupme.Attachment{{Type: "image"}},
		})
		Expect(middleUser.NumFavorites()).To(Equal(2))

		winningUser := groupme.User{
			ID:    "winning-user",
			Memes: groupme.NewMessageHeap(0),
		}
		winningUser.AddMeme(groupme.Message{
			ID:          "3",
			FavoritedBy: []string{"1", "2", "3"},
			Attachments: []groupme.Attachment{{Type: "image"}},
		})
		Expect(winningUser.NumFavorites()).To(Equal(3))

		heap.Push(&losingUser)
		Expect(heap.Len()).To(Equal(1))

		heap.Push(&winningUser)
		Expect(heap.Len()).To(Equal(2))

		heap.Push(&middleUser)
		Expect(heap.Len()).To(Equal(3))

		actualUser := heap.Pop()
		Expect(actualUser).To(Equal(winningUser))

		actualUser = heap.Pop()
		Expect(actualUser).To(Equal(middleUser))
	})

})
