package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenaivuong/litebank/internal/models"
	"github.com/nguyenaivuong/litebank/internal/services"
	"github.com/nguyenaivuong/litebank/internal/token"
)

type TransferHandlers struct {
	transferService *services.TransferService
	accountService  services.AccountService
}

func NewTransferHandlers(transferService *services.TransferService) *TransferHandlers {
	return &TransferHandlers{transferService: transferService}
}

type transferRequest struct {
	FromAccountNumber string  `json:"from_account_id" binding:"required,min=1"`
	ToAccountNumber   string  `json:"to_account_id" binding:"required,min=1"`
	Amount            float64 `json:"amount" binding:"required,gt=0"`
	Currency          string  `json:"currency" binding:"required,currency"`
}
type TransferTxParams struct {
	FromAccountNumber string  `json:"from_account_id"`
	ToAccountNumber   string  `json:"to_account_id"`
	Amount            float64 `json:"amount"`
}

func (c *TransferHandlers) TransferHandler(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := c.accountService.ValidAccount(ctx, req.FromAccountNumber, req.Currency)
	if !valid {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.User.ID != authPayload.UserID {
		err := errors.New("from account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = c.accountService.ValidAccount(ctx, req.FromAccountNumber, req.Currency)
	if !valid {
		return
	}

	arg := models.TransferTxParams{
		FromAccountNumber: req.FromAccountNumber,
		ToAccountNumber:   req.ToAccountNumber,
		Amount:            req.Amount,
	}

	result, err := c.transferService.TransferTx(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
