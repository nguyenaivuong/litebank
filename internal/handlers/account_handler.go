// internal/handlers/account_handler.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/internal/models"
	"github.com/nguyenaivuong/litebank/internal/services"
	"github.com/nguyenaivuong/litebank/internal/token"
)

// AccountHandler handles HTTP requests related to accounts
type AccountHandler struct {
	AccountService *services.AccountService
}

// NewAccountHandler creates a new instance of AccountHandler
func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetAccountByID handles the retrieval of an account by ID
func (h *AccountHandler) GetAccountByID(ctx *gin.Context) {
	accountNumber := ctx.Param("account_number")

	account, err := h.AccountService.GetAccountByID(accountNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// UpdateAccount handles the update of an existing account
func (h *AccountHandler) UpdateAccount(ctx *gin.Context) {
	accountNumber := ctx.Param("account_number")

	var account models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the account exists
	_, err := h.AccountService.GetAccountByID(accountNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if err := h.AccountService.UpdateAccount(&account); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account"})
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// DeleteAccount handles the deletion of an account by ID
func (h *AccountHandler) DeleteAccount(ctx *gin.Context) {
	accountNumber := ctx.Param("account_number")

	// Check if the account exists
	_, err := h.AccountService.GetAccountByID(accountNumber)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if err := h.AccountService.DeleteAccount(accountNumber); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete account"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListAccounts handles the retrieval of all accounts
func (h *AccountHandler) ListAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := services.ListAccountsParams{
		UserID: authPayload.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := h.AccountService.ListAccounts(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
