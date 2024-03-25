package handler

import (
	"net/http"

	"github.com/ArtuoS/payment-service-provider/internal/domain"
	"github.com/ArtuoS/payment-service-provider/internal/service"
	"github.com/gin-gonic/gin"
)

type PayableHandler struct {
	payableService *service.PayableService
}

func NewPayableHandler(payableService *service.PayableService) *PayableHandler {
	return &PayableHandler{
		payableService: payableService,
	}
}

func (p *PayableHandler) GetBalance(ctx *gin.Context) {
	values := ctx.Request.URL.Query()
	ctx.JSON(http.StatusOK, p.payableService.GetBalance(domain.NewGetBalanceModel(values.Get("status"), values.Get("card_owner"))))
}
