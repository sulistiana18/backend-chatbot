package repository

import (
	"context"
	"time"

	"backend/entity"

	"gorm.io/gorm"
)

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) entity.ConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) GetByID(ctx context.Context, id string) (*entity.Conversation, error) {
	var c entity.Conversation
	if err := r.db.WithContext(ctx).First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *conversationRepository) GetByUserID(ctx context.Context, userID string) ([]*entity.Conversation, error) {
	var list []*entity.Conversation
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("updated_at desc").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *conversationRepository) Create(ctx context.Context, c *entity.Conversation) error {
	return r.db.WithContext(ctx).Create(c).Error
}

func (r *conversationRepository) Touch(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&entity.Conversation{}).
		Where("id = ?", id).
		Update("updated_at", time.Now()).Error
}