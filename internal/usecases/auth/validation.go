package auth

import (
	"errors"
	"strings"
)

func ValidateCredentials(nickname, password string) error {
	nickname = strings.TrimSpace(nickname)
	password = strings.TrimSpace(password)

	if nickname == "" || password == "" {
		return errors.New("nickname or password must not be empty")
	}
	if len(nickname) < 3 || len(nickname) > 50 {
		return errors.New("nickname must be between 3 and 50 characters")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
