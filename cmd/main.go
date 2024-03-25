package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ArtuoS/payment-service-provider/internal/database"
	"github.com/ArtuoS/payment-service-provider/internal/domain"
	"github.com/ArtuoS/payment-service-provider/internal/repository"
	"github.com/ArtuoS/payment-service-provider/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	context := database.NewContext()
	defer context.DB.Close()

	payableRepository := repository.NewPayableRepository(context)
	payableService := service.NewPayableService(payableRepository)

	transactionRepository := repository.NewTransactionRepository(context)
	transactionService := service.NewTransactionService(transactionRepository, payableService)

	r := gin.Default()
	transactionsV1 := r.Group("/v1/transaction")
	{
		transactionsV1.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, transactionService.GetTransactions())
		})

		transactionsV1.POST("", func(ctx *gin.Context) {
			data, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			var createTransactionModel *domain.CreateTransactionModel
			json.Unmarshal(data, &createTransactionModel)
			err = transactionService.CreateTransaction(createTransactionModel)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			ctx.JSON(http.StatusOK, "Transaction created.")
		})
	}

	payablesV1 := r.Group("/v1/payable")
	{
		payablesV1.GET("", func(ctx *gin.Context) {
			values := ctx.Request.URL.Query()
			ctx.JSON(http.StatusOK, payableService.GetBalance(domain.NewGetBalanceModel(values.Get("status"), values.Get("card_owner"))))
		})
	}

	r.Run("127.0.0.1:8080")
}
