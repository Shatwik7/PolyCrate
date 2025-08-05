package userservice

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Username          string
	Email             string
	FullName          string
	ProfilePictureUrl string
	Bio               string
	Password          string
}

type UpdateUserInput struct {
	ID                uuid.UUID
	FullName          string
	ProfilePictureUrl string
	Bio               string
}

type LoginInput struct {
	Email    string
	Password string
}

type ChangePasswordInput struct {
	ID          uuid.UUID
	NewPassword string
}

type User struct {
	ID                uuid.UUID
	Username          string
	Email             string
	FullName          string
	ProfilePictureUrl string
	Bio               string
	Website           sql.NullString
	Location          sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type UserCredential struct {
	UserID       uuid.UUID
	PasswordHash string
	LastLogin    sql.NullTime
	IsActive     bool
}
