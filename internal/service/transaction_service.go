package service

import (
	"github.com/ArtuoS/payment-service-provider/internal/domain"
	"github.com/ArtuoS/payment-service-provider/internal/repository"
	"github.com/sirupsen/logrus"
)

type TransactionService struct {
	repository     *repository.TransactionRepository
	payableService *PayableService
}

func NewTransactionService(repository *repository.TransactionRepository, payableService *PayableService) *TransactionService {
	return &TransactionService{
		repository:     repository,
		payableService: payableService,
	}
}

func (t *TransactionService) CreateTransaction(createTransactionModel *domain.CreateTransactionModel) error {
	createTransactionModel.ApplyDiscount()
	transactionId, err := t.repository.CreateTransaction(createTransactionModel)
	if err != nil {
		logrus.Error(err)
		return err
	}

	_, err = t.payableService.repository.CreatePayable(domain.NewCreatePayableModel(transactionId, createTransactionModel))
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func (t *TransactionService) GetTransactions() []domain.GetTransactionResult {
	return t.repository.GetTransactions()
}
