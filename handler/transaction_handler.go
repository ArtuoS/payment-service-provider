package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArtuoS/payment-service-provider/internal/domain"
	"github.com/ArtuoS/payment-service-provider/internal/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (t *TransactionHandler) GetTransactions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.transactionService.GetTransactions())
}

func (t *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var createTransactionModel *domain.CreateTransactionModel
	json.Unmarshal(data, &createTransactionModel)
	err = t.transactionService.CreateTransaction(createTransactionModel)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Transaction created.")
}
