package adapters

import (
	"birthday_greetings/internal/core/domain"
	"fmt"
	"os"
	"strings"
	"time"
)

type FileFriendRepository struct {
	FilePath string
}

func (f *FileFriendRepository) GetAll() ([]domain.Friend, error) {
	content, err := os.ReadFile(f.FilePath)
	if err != nil {
		fmt.Printf("Failed to read file with path: %s", err.Error()) // replace with logger
		return nil, err
	}

	var friends []domain.Friend

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for lineIndex := 1; lineIndex < len(lines); lineIndex++ {
		line := lines[lineIndex]
		friendParts := strings.Split(line, ", ")
		firstName := friendParts[1]
		rawBirthDate := friendParts[2]
		emailAddress := friendParts[3]

		layout := "2006/01/02"
		birthDate, err := time.Parse(layout, rawBirthDate)
		if err != nil {
			fmt.Printf("Failed to format birthdate: %s", rawBirthDate) // replace with logger
			return nil, err
		}

		friend := domain.Friend{
			FirstName:    firstName,
			Birthdate:    birthDate,
			EmailAddress: emailAddress,
		}
		friends = append(friends, friend)
	}

	return friends, nil
}
