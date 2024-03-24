package service

import (
	"github.com/ArtuoS/payment-service-provider/internal/domain"
	"github.com/ArtuoS/payment-service-provider/internal/repository"
)

type PayableService struct {
	repository *repository.PayableRepository
}

func NewPayableService(repository *repository.PayableRepository) *PayableService {
	return &PayableService{
		repository: repository,
	}
}

func (t *PayableService) CreatePayable(createPayableModel *domain.CreatePayableModel) (int64, error) {
	return t.repository.CreatePayable(createPayableModel)
}

func (t *PayableService) GetBalance(createPayableModel *domain.GetBalanceModel) []domain.GetBalanceResult {
	return t.repository.GetBalance(createPayableModel)
}
