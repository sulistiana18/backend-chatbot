package handler

import (
	"net/http"

	"backend/entity"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase entity.UserUsecase
}

func NewUserHandler(r *gin.Engine, u entity.UserUsecase) {
	h := &UserHandler{usecase: u}

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.GET("/users/:id", h.GetUser)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.CreateUser(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email atau password salah"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.usecase.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}