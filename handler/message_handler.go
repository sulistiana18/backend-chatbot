package handler

import (
	"net/http"

	"backend/entity"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	usecase entity.MessageUsecase
}

func NewMessageHandler(r *gin.Engine, u entity.MessageUsecase) {
	h := &MessageHandler{usecase: u}

	r.POST("/conversations/:id/messages", h.Send)
	r.GET("/conversations/:id/messages", h.History)
	r.GET("/users/:id/messages", h.UserHistory)
}

func (h *MessageHandler) Send(c *gin.Context) {
	conversationID := c.Param("id")

	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reply, err := h.usecase.SendMessage(c.Request.Context(), conversationID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reply)
}

func (h *MessageHandler) History(c *gin.Context) {
	conversationID := c.Param("id")

	list, err := h.usecase.GetHistory(c.Request.Context(), conversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *MessageHandler) UserHistory(c *gin.Context) {
	userID := c.Param("id")

	list, err := h.usecase.GetUserHistory(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}