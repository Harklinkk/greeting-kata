package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFriendRepository struct {
	mock.Mock
}

type MockMessageSender struct {
	mock.Mock
}

func (m *MockMessageSender) Send(content string, contact string) {
	m.Called(content, contact)
}

func (m *MockFriendRepository) GetAll() ([]Friend, error) {
	args := m.Called()
	return args.Get(0).([]Friend), args.Error(1)
}

func Test_should_send_birthday_greeting_message_to_John(t *testing.T) {
	todaydate := time.Date(1969, time.October, 19, 0, 0, 0, 0, time.UTC)
	email := "John@john.fr"
	firstName := "John"
	emailAddress := "John@john.fr"
	birthdate := time.Date(2000, time.October, 19, 14, 30, 0, 0, time.UTC)
	message := "Subject: Happy birthday!\n\nHappy birthday, dear John!"
	mockMessageSender := new(MockMessageSender)
	mockFriendRepository := new(MockFriendRepository)
	mockMessageSender.On("Send", message, email)
	mockFriendRepository.On("GetAll").Return([]Friend{
		{firstName, birthdate, emailAddress},
	}, nil)
	birthdayGreeter := &BirthdayGreeter{
		Sender:           mockMessageSender,
		FriendRepository: mockFriendRepository,
	}
	birthdayGreeter.Greet(todaydate)
	mockMessageSender.AssertExpectations(t)
}

func Test_check_birthdate(t *testing.T) {
	todaydate := time.Date(1969, time.May, 05, 0, 0, 0, 0, time.UTC)
	friend := Friend{
		FirstName:    "Rachel",
		Birthdate:    time.Date(1969, time.May, 05, 0, 0, 0, 0, time.UTC),
		EmailAddress: "rgreen@scaleway.com",
	}
	assert.Equal(t, true, friend.CheckBirthdate(todaydate))
}
