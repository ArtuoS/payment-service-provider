package domain

import (
	"time"
)

type Payable struct {
	ID            int64     `db:"id" json:"id"`
	Status        Status    `db:"status" json:"status"`
	PaymentDate   time.Time `db:"payment_date" json:"payment_date"`
	TransactionId int64     `db:"transaction_id" json:"transaction_id"`
}

type CreatePayableModel struct {
	Status        Status    `db:"status" json:"status"`
	PaymentDate   time.Time `db:"payment_date" json:"payment_date"`
	TransactionId int64     `db:"transaction_id" json:"transaction_id"`
}

func NewCreatePayableModel(transactionId int64, createTransactionModel *CreateTransactionModel) *CreatePayableModel {
	status, paymentDate := func() (Status, time.Time) {
		switch createTransactionModel.PaymentMethod {
		case CreditCard:
			return Paid, createTransactionModel.CardExpirationDate.AddDate(0, 0, 30)
		case DebitCard:
			return WaitingFunds, createTransactionModel.CardExpirationDate
		}
		return WaitingFunds, createTransactionModel.CardExpirationDate
	}()

	return &CreatePayableModel{
		Status:        status,
		PaymentDate:   paymentDate,
		TransactionId: transactionId,
	}
}

type Status uint8

const (
	Paid Status = iota
	WaitingFunds
)

func (s Status) Sting() string {
	switch s {
	case Paid:
		return "paid"
	case WaitingFunds:
		return "waiting_funds"
	}
	return "unknown"
}
