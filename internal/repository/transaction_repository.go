package repository

import (
	"log"

	"github.com/ArtuoS/payment-service-provider/internal/database"
	"github.com/ArtuoS/payment-service-provider/internal/domain"
)

type TransactionRepository struct {
	Context *database.Context
}

func NewTransactionRepository(context *database.Context) *TransactionRepository {
	return &TransactionRepository{
		Context: context,
	}
}

func (t *TransactionRepository) CreateTransaction(createTransactionModel *domain.CreateTransactionModel) (int64, error) {
	var lastInsertedId int64
	query := `INSERT INTO transactions 
               (transaction_value, description, payment_method, card_number, card_owner, card_expiration_date, card_cvv) 
               VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err := t.Context.DB.QueryRow(query, createTransactionModel.Value, createTransactionModel.Description, createTransactionModel.PaymentMethod,
		createTransactionModel.CardNumber, createTransactionModel.CardOwner, createTransactionModel.CardExpirationDate, createTransactionModel.CardCvv).Scan(&lastInsertedId)
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

func (t *TransactionRepository) GetTransactions() []domain.Transaction {
	transactions := []domain.Transaction{}
	transaction := domain.Transaction{}
	rows, _ := t.Context.DB.Queryx("SELECT ID, transaction_value, description, payment_method, card_number, card_owner, card_expiration_date, card_cvv FROM transactions")
	for rows.Next() {
		err := rows.StructScan(&transaction)
		if err != nil {
			log.Fatalln(err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}
