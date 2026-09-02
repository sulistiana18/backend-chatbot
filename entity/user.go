package entity

import (
	"context"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
}

type UserUsecase interface {
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, name, email, password string) (*User, error)
	Login(ctx context.Context, email, password string) (*User, error)
}