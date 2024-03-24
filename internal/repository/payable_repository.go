package repository

import (
	"fmt"
	"log"

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

func (t *PayableRepository) GetBalance(createPayableModel *domain.GetBalanceModel) []domain.GetBalanceResult {
	payables := []domain.GetBalanceResult{}
	payable := domain.GetBalanceResult{}
	rows, _ := t.Context.DB.Queryx("select p.id, p.status, p.payment_date, t.transaction_value, t.card_owner from payables p inner join transactions t on t.id = p.transaction_id where p.status = $1 and t.card_owner = $2", createPayableModel.Status, createPayableModel.CardOwner)
	for rows.Next() {
		err := rows.StructScan(&payable)
		if err != nil {
			log.Fatalln(err)
		}

		payable.Format()
		payables = append(payables, payable)
	}

	return payables
}
