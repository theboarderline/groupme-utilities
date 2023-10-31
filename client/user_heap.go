package groupme

import "container/heap"

type UserHeap struct {
	capacity int
	users    []User
}

func GetTopNUsers(users []User, n int) []User {

	userHeap := NewUserHeap(n)

	for _, user := range users {
		heap.Push(userHeap, user)
		if userHeap.Len() > n {
			heap.Pop(userHeap)
		}
	}

	return userHeap.List()
}

func NewUserHeap(capacity int) *UserHeap {
	h := UserHeap{
		capacity: capacity,
		users:    make([]User, 0, capacity),
	}

	heap.Init(&h)
	return &h
}

func (h *UserHeap) Len() int {
	return len(h.users)
}

func (h *UserHeap) Less(i, j int) bool {
	return h.users[i].NumFavorites() > h.users[j].NumFavorites()
}

func (h *UserHeap) Swap(i, j int) {
	h.users[i], h.users[j] = h.users[j], h.users[i]
}

func (h *UserHeap) IsFull() bool {
	return h.capacity != 0 && len(h.users) >= h.capacity
}

func (h *UserHeap) UserHasMoreLikes(userOne, userTwo User) bool {
	return userOne.NumFavorites() > userTwo.NumFavorites()
}

func (h *UserHeap) Push(x interface{}) {

	user, ok := x.(*User)
	if !ok {
		return
	}

	if h.IsFull() {
		if h.UserHasMoreLikes(*user, h.users[0]) {
			heap.Pop(h)
		} else {
			return
		}
	}

	h.users = append(h.users, *user)
	heap.Fix(h, h.Len()-1)
}

func (h *UserHeap) Pop() interface{} {
	old := h.users
	n := len(old)
	x := old[0]
	old[0] = old[n-1]
	h.users = old[0 : n-1]
	heap.Fix(h, 0)
	return x
}

func (h *UserHeap) List() (users []User) {
	users = make([]User, len(h.users))
	copy(users, h.users)

	for i, j := 0, len(users)-1; i < j; i, j = i+1, j-1 {
		users[i], users[j] = users[j], users[i]
	}

	return users
}
