package auth

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
	"github.com/fleischgewehr/crypto-labs/passwords/internal/lib"
)

const (
	KeyTimespan = 3 * 60 * 60 // 3 hrs in seconds
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

func HandleInvalidPassword(app *app.Application, username string) error {
	key := fmt.Sprintf("invalid-login:%s", username)
	currentCount, err := app.Cache.Client.Get(key)
	if err != nil {
		app.Cache.Client.Set(&memcache.Item{Key: key, Value: lib.IntToBytes(1), Expiration: KeyTimespan})
		return nil
	}

	countValue := lib.BytesToInt(currentCount.Value)
	if countValue > 20 {
		return errors.New("you have exceeded max number of login attempts, try again later")
	} else {
		app.Cache.Client.Set(&memcache.Item{Key: key, Value: lib.IntToBytes(countValue + 1), Expiration: KeyTimespan})
		return nil
	}
}
