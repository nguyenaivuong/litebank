package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/internal/models"
	"github.com/nguyenaivuong/litebank/internal/services"
	"github.com/nguyenaivuong/litebank/internal/utils"
)

type SignupHandler struct {
	signupService services.SignupService
}

func NewSignupHandler(signupService services.SignupService) *SignupHandler {
	return &SignupHandler{signupService: signupService}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum" error_message:"Username is required and must contain only alphanumeric characters"`
	Password string `json:"password" binding:"required" error_message:"Password is required"`
	Email    string `json:"email" binding:"omitempty,email" error_message:"Invalid email format"`
	FullName string `json:"full_name" binding:"required" error_message:"Full name is required"`
}

func (c *SignupHandler) SignupHandler(ctx *gin.Context) {
	// Extract user information from request
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.LogWithCallerInfo(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		utils.LogWithCallerInfo(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user := &models.User{
		Username:  req.Username,
		Password:  hashedPassword,
		Email:     req.Email,
		FullName:  req.FullName,
		CreatedAt: time.Now(),
	}

	// Signup user
	err = c.signupService.Signup(ctx, user)
	if err != nil {
		utils.LogWithCallerInfo("")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	utils.LogWithCallerInfo("Create user and acccount successful.")

	// Respond with success message
	ctx.JSON(http.StatusCreated, gin.H{"message": "Signup successful."})
}

func hashPassword(password string) (string, error) {
	// Implement secure password hashing algorithm
	// This is just a placeholder
	return password, nil
}
