package main

import (
	"log"

	"backend/ai"
	"backend/entity"
	"backend/handler"
	"backend/repository"
	"backend/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres1 dbname=chatbot_test1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate database tables
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Conversation{},
		&entity.Message{},
	); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	r := gin.Default()

	// --- Dependency Injection ---

	// User
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(r, userUsecase)

	// Conversation
	conversationRepo := repository.NewConversationRepository(db)
	conversationUsecase := usecase.NewConversationUsecase(conversationRepo)
	handler.NewConversationHandler(r, conversationUsecase)

	// Message (butuh conversationRepo untuk Touch, dan aiProvider untuk balasan)
	messageRepo := repository.NewMessageRepository(db)
	aiProvider := ai.NewDummyProvider() // nanti tinggal ganti ai.NewOpenAIProvider(apiKey)
	messageUsecase := usecase.NewMessageUsecase(messageRepo, conversationRepo, aiProvider)
	handler.NewMessageHandler(r, messageUsecase)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}