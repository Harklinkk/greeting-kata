package domain

import "time"

type Greeter interface {
	Greet(todayDate time.Time)
}
