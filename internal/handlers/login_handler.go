package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/internal/services"
	"github.com/nguyenaivuong/litebank/internal/utils"
)

type LoginHandler struct {
	loginService services.LoginService
}

func NewLoginHandler(loginService services.LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum" error_message:"Username is required and must contain only alphanumeric characters"`
	Password string `json:"password" binding:"required" error_message:"Password is required"`
}

func (c *LoginHandler) LoginHandler(ctx *gin.Context) {
	// Extract user information from request
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.LogWithCallerInfo(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login user
	token, err := c.loginService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"message": "Invalid username or password."})
		return
	}

	// Return success response with token
	ctx.JSON(200, gin.H{"message": "Login successful.", "token": token})
}
