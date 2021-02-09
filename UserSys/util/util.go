package util

import (
	"regexp"
)

const (
	str = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

func PasswordCheck(password string) bool {
	if len(password) >= 5 {
		return true
	} else {
		return false
	}
}

func EmailCheck(email string) bool {
	check := regexp.MustCompile(str)

	if check.MatchString(email) {
		//valid
		return true
	} else{
		return false
	}
}
