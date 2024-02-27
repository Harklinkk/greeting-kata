package domain

import (
	"fmt"
	"time"
)

type BirthdayGreeter struct {
	Sender           MessageSender
	FriendRepository FriendRepository
}

func (b *BirthdayGreeter) Greet(todayDate time.Time) {
	friends, _ := b.FriendRepository.GetAll()
	for _, friend := range friends {
		if friend.CheckBirthdate(todayDate) {
			content := fmt.Sprintf("Subject: Happy birthday!\n\nHappy birthday, dear %s!", friend.FirstName)
			b.Sender.Send(content, friend.EmailAddress)
		}
	}

}
