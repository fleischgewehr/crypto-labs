package models

import (
	"context"

	"github.com/fleischgewehr/crypto-labs/passwords/internal/app"
)

type User struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	PasswordHash      string `json:"password_hash"`
	PasswordSalt      string `json:"password_salt"`
	InvalidLoginCount int    `json:"invalid_login_count"`
	IsBlocked         bool   `json:"is_blocked"`
}

func (u *User) Create(ctx context.Context, app *app.Application) error {
	stmt := `
		insert into users(username, password_hash, password_salt)
		values(?,?,?)
		returning id
	`

	return app.DB.Client.
		QueryRowContext(ctx, stmt, u.Username, u.PasswordHash, u.PasswordSalt).
		Scan(&u.ID)
}

func (u *User) GetById(ctx context.Context, app *app.Application) error {
	stmt := "select * from users where id = ?"

	return app.DB.Client.
		QueryRowContext(ctx, stmt, u.ID).
		Scan(&u.ID, &u.Username, &u.PasswordHash, &u.PasswordSalt)
}

func (u *User) GetByUsername(ctx context.Context, app *app.Application) error {
	stmt := "select * from users where username = ?"

	return app.DB.Client.
		QueryRowContext(ctx, stmt, u.Username).
		Scan(&u.ID, &u.Username, &u.PasswordHash, &u.PasswordSalt)
}

func (u *User) BlockAccount(ctx context.Context, app *app.Application) {
	stmt := "update users set is_blocked = true where id = ?"
	app.DB.Client.QueryRowContext(ctx, stmt, u.ID)
}

func (u *User) UpsertInvalidLoginCount(ctx context.Context, app *app.Application) {
	stmt := `
		insert into users(username, invalid_login_count)
		values(?,?)
		on conflict (username) do update
		set invalid_login_count = EXCLUDED.invalid_login_count + 1
	`
	app.DB.Client.QueryRowContext(ctx, stmt, u.Username, u.InvalidLoginCount+1)
}
