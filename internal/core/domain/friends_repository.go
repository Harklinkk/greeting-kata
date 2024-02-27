package domain

type FriendRepository interface {
	GetAll() ([]Friend, error)
}
