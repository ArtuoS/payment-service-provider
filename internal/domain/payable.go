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

type GetBalanceModel struct {
	CardOwner string `db:"card_owner" json:"card_owner"`
	Status    Status `db:"status" json:"status"`
}

type GetBalanceResult struct {
	ID          int64     `db:"id" json:"id"`
	StatusDb    Status    `db:"status" json:"-"`
	Status      string    `json:"status"`
	CardOwner   string    `db:"card_owner" json:"card_owner"`
	PaymentDate time.Time `db:"payment_date" json:"payment_date"`
	Value       float32   `db:"transaction_value" json:"transaction_value"`
}

func (g *GetBalanceResult) Format() {
	g.Status = g.StatusDb.String()
}

func NewGetBalanceModel(status string, cardOwner string) *GetBalanceModel {
	return &GetBalanceModel{
		Status: func() Status {
			switch status {
			case "0":
				return Paid
			case "1":
				return WaitingFunds
			}
			return WaitingFunds
		}(),
		CardOwner: cardOwner,
	}
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

func (s Status) String() string {
	switch s {
	case Paid:
		return "paid"
	case WaitingFunds:
		return "waiting_funds"
	}
	return "unknown"
}
