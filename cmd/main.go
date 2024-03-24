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
	v1 := r.Group("/v1")
	{
		v1.GET("/transaction", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, transactionService.GetTransactions())
		})

		v1.POST("/transaction", func(ctx *gin.Context) {
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

	r.Run("127.0.0.1:8080")
}
