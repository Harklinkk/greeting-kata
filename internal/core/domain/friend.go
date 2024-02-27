package domain

import "time"

type Friend struct {
	FirstName    string
	Birthdate    time.Time
	EmailAddress string
}

func (f *Friend) CheckBirthdate(todaydate time.Time) bool {
	return (f.Birthdate.Day() == todaydate.Day() && f.Birthdate.Month() == todaydate.Month())
}
