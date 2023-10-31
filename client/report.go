package groupme

import (
	"fmt"
	"time"
)

type Report struct {
	StartDate *time.Time
	EndDate   *time.Time
	Users     map[string]*User
	UsersHeap *UserHeap
	MemesHeap *MessageHeap
}

func NewReport(memes []Message) *Report {
	report := Report{
		Users:     make(map[string]*User),
		UsersHeap: NewUserHeap(0),
		MemesHeap: NewMessageHeap(0),
	}

	for _, meme := range memes {
		report.AddMessage(meme)
	}

	return &report
}

func (r *Report) PopMeme() Message {
	m := r.MemesHeap.Pop()
	return m.(Message)
}

func (r *Report) PopUser() User {
	return r.UsersHeap.Pop().(User)
}

func (r *Report) AddMessage(message Message) {
	r.MemesHeap.Push(message)
	r.AddUser(NewUser(message))
}

func (r *Report) AddUser(newUser *User) {
	existingUser, ok := r.Users[newUser.ID]
	if !ok {
		r.Users[newUser.ID] = newUser
		r.UsersHeap.Push(newUser)
	} else {
		for _, meme := range newUser.Memes.List() {
			existingUser.Memes.Push(meme)
		}
	}
}

func (r *Report) String() string {
	line1 := "Meme Report:\n\n"

	line2 := fmt.Sprintf("Number of memes: %d\n", r.MemesHeap.Len())
	line3 := fmt.Sprintf("Number of favorites: %d\n", r.MemesHeap.NumFavorites())
	return line1 + line2 + line3
}
