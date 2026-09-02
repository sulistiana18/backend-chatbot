package usecase

import (
	"context"

	"backend/entity"

	"github.com/google/uuid"
)

type messageUsecase struct {
	messageRepo      entity.MessageRepository
	conversationRepo entity.ConversationRepository
	aiProvider       entity.AIProvider
}

func NewMessageUsecase(
	messageRepo entity.MessageRepository,
	conversationRepo entity.ConversationRepository,
	aiProvider entity.AIProvider,
) entity.MessageUsecase {
	return &messageUsecase{
		messageRepo:      messageRepo,
		conversationRepo: conversationRepo,
		aiProvider:       aiProvider,
	}
}

func (u *messageUsecase) SendMessage(ctx context.Context, conversationID, content string) (*entity.Message, error) {
	// 1. simpan pesan dari user
	userMsg := &entity.Message{
		ID:             uuid.NewString(),
		ConversationID: conversationID,
		Role:           entity.RoleUser,
		Content:        content,
	}
	if err := u.messageRepo.Create(ctx, userMsg); err != nil {
		return nil, err
	}
	_ = u.conversationRepo.Touch(ctx, conversationID)

	// 2. ambil histori lengkap buat konteks AI
	history, err := u.messageRepo.GetByConversationID(ctx, conversationID)
	if err != nil {
		return nil, err
	}
	historyValues := make([]entity.Message, len(history))
	for i, m := range history {
		historyValues[i] = *m
	}

	// 3. generate balasan dari AI provider
	reply, err := u.aiProvider.GenerateReply(ctx, historyValues, content)
	if err != nil {
		return nil, err
	}

	// 4. simpan balasan AI sebagai Message baru
	aiMsg := &entity.Message{
		ID:             uuid.NewString(),
		ConversationID: conversationID,
		Role:           entity.RoleAssistant,
		Content:        reply.Content,
		Options:        reply.Options,
	}
	if err := u.messageRepo.Create(ctx, aiMsg); err != nil {
		return nil, err
	}
	_ = u.conversationRepo.Touch(ctx, conversationID)

	return aiMsg, nil
}

func (u *messageUsecase) GetHistory(ctx context.Context, conversationID string) ([]*entity.Message, error) {
	return u.messageRepo.GetByConversationID(ctx, conversationID)
}

func (u *messageUsecase) GetUserHistory(ctx context.Context, userID string) ([]*entity.Message, error) {
	return u.messageRepo.GetByUserID(ctx, userID)
}