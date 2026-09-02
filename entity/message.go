package entity

import (
	"context"
	"time"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type Message struct {
	ID             string    `gorm:"primaryKey;type:uuid" json:"id"`
	ConversationID string    `gorm:"not null;index:idx_conv_created" json:"conversation_id"`
	Role           Role      `gorm:"not null" json:"role"`
	Content        string    `gorm:"type:text;not null" json:"content"`
	Options        []string  `gorm:"type:jsonb" json:"options,omitempty"`
	CreatedAt      time.Time `gorm:"index:idx_conv_created" json:"created_at"`
}

type MessageRepository interface {
	Create(ctx context.Context, m *Message) error
	GetByConversationID(ctx context.Context, conversationID string) ([]*Message, error)
	GetByUserID(ctx context.Context, userID string) ([]*Message, error)
}

type MessageUsecase interface {
	SendMessage(ctx context.Context, conversationID, content string) (*Message, error)
	GetHistory(ctx context.Context, conversationID string) ([]*Message, error)
	GetUserHistory(ctx context.Context, userID string) ([]*Message, error)
}