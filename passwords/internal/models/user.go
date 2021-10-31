package models

import (
	"context"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`
}

func (u *User) Create(ctx context.Context, app *app.Application) error {
	stmt := `
		insert into users(username, password_hash, password_salt)
		values(?,?,?)
		returning id
	`
	err := app.DB.Client.
		QueryRowContext(ctx, stmt, u.Username, u.PasswordHash, u.PasswordSalt).
		Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetById(ctx context.Context, app *app.Application) error {
	stmt := "select * from users where id = ?"
	err := app.DB.Client.
		QueryRowContext(ctx, stmt, u.ID).
		Scan(&u.ID, &u.Username, &u.PasswordHash, &u.PasswordSalt)
	if err != nil {
		return err
	}

	return nil
}
