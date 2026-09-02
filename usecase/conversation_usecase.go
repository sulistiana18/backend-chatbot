package usecase

import (
	"context"

	"backend/entity"

	"github.com/google/uuid"
)

type conversationUsecase struct {
	conversationRepo entity.ConversationRepository
}

func NewConversationUsecase(repo entity.ConversationRepository) entity.ConversationUsecase {
	return &conversationUsecase{conversationRepo: repo}
}

func (u *conversationUsecase) StartConversation(ctx context.Context, userID, title string) (*entity.Conversation, error) {
	c := &entity.Conversation{
		ID:     uuid.NewString(),
		UserID: userID,
		Title:  title,
	}
	if err := u.conversationRepo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (u *conversationUsecase) ListConversations(ctx context.Context, userID string) ([]*entity.Conversation, error) {
	return u.conversationRepo.GetByUserID(ctx, userID)
}