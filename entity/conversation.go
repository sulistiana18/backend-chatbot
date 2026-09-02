package entity

import (
	"context"
	"time"
)

type Conversation struct {
	ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
	UserID    string    `gorm:"not null;index:idx_user_updated" json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"index:idx_user_updated" json:"updated_at"`
}

type ConversationRepository interface {
	GetByID(ctx context.Context, id string) (*Conversation, error)
	GetByUserID(ctx context.Context, userID string) ([]*Conversation, error)
	Create(ctx context.Context, c *Conversation) error
	Touch(ctx context.Context, id string) error
}

type ConversationUsecase interface {
	StartConversation(ctx context.Context, userID, title string) (*Conversation, error)
	ListConversations(ctx context.Context, userID string) ([]*Conversation, error)
}