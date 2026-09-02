package repository

import (
	"context"

	"backend/entity"

	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) entity.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(ctx context.Context, m *entity.Message) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *messageRepository) GetByConversationID(ctx context.Context, conversationID string) ([]*entity.Message, error) {
	var list []*entity.Message
	err := r.db.WithContext(ctx).
		Where("conversation_id = ?", conversationID).
		Order("created_at asc").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *messageRepository) GetByUserID(ctx context.Context, userID string) ([]*entity.Message, error) {
	var list []*entity.Message
	err := r.db.WithContext(ctx).
		Joins("JOIN conversations ON conversations.id = messages.conversation_id").
		Where("conversations.user_id = ?", userID).
		Order("messages.created_at asc").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}