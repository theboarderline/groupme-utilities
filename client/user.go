package groupme

type User struct {
	ID    string
	Name  string
	Memes *MessageHeap
}

func NewUser(message Message) *User {
	user := User{
		ID:    message.SenderID,
		Name:  message.Name,
		Memes: NewMessageHeap(0),
	}
	user.AddMeme(message)
	return &user
}

func (u User) NumMemes() int {
	return u.Memes.Len()
}

func (u User) NumFavorites() int {
	return u.Memes.NumFavorites()
}

func (u User) AddMeme(m Message) {
	if m.IsMeme() {
		u.Memes.Push(m)
	}
}
