package adapters

import (
	"birthday_greetings/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Should_Return_Friend_List_When_Get_All(t *testing.T) {
	filePath := "../../tests/resources/friend_list.txt"
	expected := []domain.Friend{
		{
			FirstName:    "Christophe",
			Birthdate:    time.Date(1995, time.October, 31, 0, 0, 0, 0, time.UTC),
			EmailAddress: "cmach@gmail.com",
		},
		{
			FirstName:    "Mary",
			Birthdate:    time.Date(1975, time.September, 11, 0, 0, 0, 0, time.UTC),
			EmailAddress: "mary.ann@foobar.com",
		},
	}
	friendRepository := FileFriendRepository{
		FilePath: filePath,
	}

	friends, _ := friendRepository.GetAll()

	assert.Equal(t, expected, friends)
}
