// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          int64        `json:"id"`
	Userid      int64        `json:"userid"`
	Email       string       `json:"email"`
	Amount      string       `json:"amount"`
	Description string       `json:"description"`
	Tag         string       `json:"tag"`
	Date        time.Time    `json:"date"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Userid       int64     `json:"userid"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID        int64        `json:"id"`
	Firstname string       `json:"firstname"`
	Lastname  string       `json:"lastname"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Phone     string       `json:"phone"`
	UserType  string       `json:"user_type"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
