package handler

import (
	"net/http"

	"backend/entity"

	"github.com/gin-gonic/gin"
)

type ConversationHandler struct {
	usecase entity.ConversationUsecase
}

func NewConversationHandler(r *gin.Engine, u entity.ConversationUsecase) {
	h := &ConversationHandler{usecase: u}

	group := r.Group("/conversations")
	group.POST("", h.Create)
	group.GET("", h.List)
}

func (h *ConversationHandler) Create(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
		Title  string `json:"title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conv, err := h.usecase.StartConversation(c.Request.Context(), req.UserID, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, conv)
}

func (h *ConversationHandler) List(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	list, err := h.usecase.ListConversations(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}