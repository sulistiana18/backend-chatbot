package entity

import "context"

type AIReply struct {
	Content string
	Options []string
}

type AIProvider interface {
	GenerateReply(ctx context.Context, history []Message, newMessage string) (*AIReply, error)
}