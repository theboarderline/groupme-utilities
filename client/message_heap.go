package groupme

import (
	"container/heap"
	"sort"
)

type MessageHeap struct {
	capacity       int
	totalFavorites int
	messages       []Message
}

func GetTopNMessages(messages []Message, n int) []Message {

	memeHeap := NewMessageHeap(n)

	for _, message := range messages {
		heap.Push(memeHeap, message)
		if memeHeap.Len() > n {
			heap.Pop(memeHeap)
		}
	}

	return memeHeap.List()
}

func NewMessageHeap(capacity int) *MessageHeap {
	h := MessageHeap{
		capacity: capacity,
		messages: make([]Message, 0, capacity),
	}

	heap.Init(&h)
	return &h
}

func (h *MessageHeap) NumFavorites() int {
	return h.totalFavorites
}

func (h *MessageHeap) Len() int {
	return len(h.messages)
}

func (h *MessageHeap) Less(i, j int) bool {
	return h.messages[i].NumFavorites() < h.messages[j].NumFavorites()
}

func (h *MessageHeap) Swap(i, j int) {
	h.messages[i], h.messages[j] = h.messages[j], h.messages[i]
}

func (h *MessageHeap) IsFull() bool {
	return h.capacity != 0 && len(h.messages) >= h.capacity
}

func (h *MessageHeap) MessageHasMoreLikes(messageOne, messageTwo Message) bool {
	return messageOne.NumFavorites() > messageTwo.NumFavorites()
}

func (h *MessageHeap) Push(x interface{}) {

	message, ok := x.(Message)
	if !ok {
		return
	}

	h.totalFavorites += message.NumFavorites()

	if h.IsFull() {
		if h.MessageHasMoreLikes(message, h.messages[0]) {
			heap.Pop(h)
		} else {
			return
		}
	}

	h.messages = append(h.messages, message)
	heap.Fix(h, h.Len()-1)
}

func (h *MessageHeap) Pop() interface{} {
	old := h.messages
	n := len(old)
	x := old[n-1]
	h.messages = old[0 : n-1]
	heap.Fix(h, 0)
	return x
}

func (h *MessageHeap) List() []Message {
	messages := make([]Message, len(h.messages))
	copy(messages, h.messages)

	sort.Slice(messages, func(i, j int) bool {
		return len(messages[i].FavoritedBy) > len(messages[j].FavoritedBy)
	})

	return messages
}
