package main

import (
	"github.com/ArtuoS/payment-service-provider/handler"
	"github.com/ArtuoS/payment-service-provider/internal/database"
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

	payableHandler := handler.NewPayableHandler(payableService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()
	transactionsV1 := r.Group("/v1/transaction")
	{
		transactionsV1.GET("", transactionHandler.GetTransactions)
		transactionsV1.POST("", transactionHandler.CreateTransaction)
	}

	payablesV1 := r.Group("/v1/payable")
	{
		payablesV1.GET("", payableHandler.GetBalance)
	}

	r.Run("127.0.0.1:8080")
}
