package user

import (
	"fmt"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(value string) (Email, error) {
	if !emailRegex.MatchString(value) {
		return Email{}, fmt.Errorf("invalid email format")
	}
	return Email{value: value}, nil
}

func (e Email) String() string {
	return e.value
}
