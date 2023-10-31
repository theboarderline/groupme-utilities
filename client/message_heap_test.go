package groupme_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/theboarderline/groupme-utilities/client"
)

var _ = Describe("MessageHeap", func() {

	It("can sort two messages and return the one with most likes", func() {
		heap := groupme.NewMessageHeap(2)
		Expect(heap).NotTo(BeNil())

		losingMessage := groupme.Message{
			FavoritedBy: []string{
				"123",
			},
		}
		middleMessage := groupme.Message{
			FavoritedBy: []string{
				"123",
				"321",
			},
		}
		winningMessage := groupme.Message{
			FavoritedBy: []string{
				"123",
				"456",
				"234",
			},
		}

		heap.Push(losingMessage)
		Expect(heap.Len()).To(Equal(1))

		heap.Push(winningMessage)
		Expect(heap.Len()).To(Equal(2))

		heap.Push(middleMessage)
		Expect(heap.Len()).To(Equal(2))

		actualMessage := heap.Pop()
		Expect(actualMessage).To(Equal(winningMessage))

		actualMessage = heap.Pop()
		Expect(actualMessage).To(Equal(middleMessage))
	})

})
