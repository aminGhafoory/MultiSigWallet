// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    user_id,
    created_at,
    updated_at,
    email,
    password_hash
    )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    )
RETURNING user_id, created_at, updated_at, email, password_hash
`

type CreateUserParams struct {
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Email,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
	)
	return i, err
}

const deleteUserSession = `-- name: DeleteUserSession :exec
DELETE 
    FROM sessions 
    WHERE token_hash=$1
`

func (q *Queries) DeleteUserSession(ctx context.Context, tokenHash string) error {
	_, err := q.db.ExecContext(ctx, deleteUserSession, tokenHash)
	return err
}

const userByEmail = `-- name: UserByEmail :one
SELECT
    user_id,
    password_hash
FROM
    users
WHERE email =$1
`

type UserByEmailRow struct {
	UserID       uuid.UUID
	PasswordHash string
}

func (q *Queries) UserByEmail(ctx context.Context, email string) (UserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, userByEmail, email)
	var i UserByEmailRow
	err := row.Scan(&i.UserID, &i.PasswordHash)
	return i, err
}

const userBySession = `-- name: UserBySession :one
SELECT users.user_id,
       email,
       password_hash
	FROM sessions
	INNER JOIN users ON
		sessions.user_id = users.user_id
	WHERE token_hash = $1
`

type UserBySessionRow struct {
	UserID       uuid.UUID
	Email        string
	PasswordHash string
}

func (q *Queries) UserBySession(ctx context.Context, tokenHash string) (UserBySessionRow, error) {
	row := q.db.QueryRowContext(ctx, userBySession, tokenHash)
	var i UserBySessionRow
	err := row.Scan(&i.UserID, &i.Email, &i.PasswordHash)
	return i, err
}
