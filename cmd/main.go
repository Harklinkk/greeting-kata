package main

import (
	"birthday_greetings/internal/adapters"
	"birthday_greetings/internal/core/domain"
	"time"
)

func main() {
	friendRepository := &adapters.FileFriendRepository{
		FilePath: "tests/resources/friend_list.txt",
	}
	messageSender := &adapters.EmailMessageSender{}

	birthdayGreeter := domain.BirthdayGreeter{
		Sender:           messageSender,
		FriendRepository: friendRepository,
	}

	birthdayGreeter.Greet(time.Now())
}
