package repository

import (
	"fmt"

	"github.com/ArtuoS/payment-service-provider/internal/database"
	"github.com/ArtuoS/payment-service-provider/internal/domain"
)

type PayableRepository struct {
	Context *database.Context
}

func NewPayableRepository(context *database.Context) *PayableRepository {
	return &PayableRepository{
		Context: context,
	}
}

func (p *PayableRepository) CreatePayable(createPayableModel *domain.CreatePayableModel) (int64, error) {
	fmt.Println(createPayableModel)

	var lastInsertedId int64
	query := `INSERT INTO payables (status, payment_date, transaction_id) VALUES ($1, $2, $3) RETURNING id`
	err := p.Context.DB.QueryRow(query, createPayableModel.Status, createPayableModel.PaymentDate, createPayableModel.TransactionId).Scan(&lastInsertedId)
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}
