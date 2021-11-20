package auth

import (
	"errors"
	"unicode"
)

func verifyPassword(password string) (upper, lower, punct, digit bool) {
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsPunct(c):
			punct = true
		case unicode.IsDigit(c):
			digit = true
		}
	}

	return upper, lower, punct, digit
}

func validatePassword(password string) error {
	upper, lower, punct, digit := verifyPassword(password)
	if !upper {
		return errors.New("password must contain at least one uppercase letter")
	} else if !lower {
		return errors.New("password must contain at least one lowercase letter")
	} else if !punct {
		return errors.New("password must contain at least one special character")
	} else if !digit {
		return errors.New("password must contain at least one digit")
	} else if len(password) < 10 {
		return errors.New("password must be longer or equal to 10 characters")
	}

	return nil
}
