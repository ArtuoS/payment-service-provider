package domain

import (
	"time"
)

type Transaction struct {
	ID                 int64         `db:"id" json:"id"`
	Value              float32       `db:"transaction_value" json:"transaction_value"`
	Description        string        `db:"description" json:"description"`
	PaymentMethod      PaymentMethod `db:"payment_method" json:"payment_method"`
	CardNumber         string        `db:"card_number" json:"card_number"`
	CardOwner          string        `db:"card_owner" json:"card_owner"`
	CardExpirationDate time.Time     `db:"card_expiration_date" json:"card_expiration_date"`
	CardCvv            uint16        `db:"card_cvv" json:"card_cvv"`
}

type CreateTransactionModel struct {
	Value              float32       `db:"transaction_value" json:"transaction_value"`
	Description        string        `db:"description" json:"description"`
	PaymentMethod      PaymentMethod `db:"payment_method" json:"payment_method"`
	CardNumber         string        `db:"card_number" json:"card_number"`
	CardOwner          string        `db:"card_owner" json:"card_owner"`
	CardExpirationDate time.Time     `db:"card_expiration_date" json:"card_expiration_date"`
	CardCvv            uint16        `db:"card_cvv" json:"card_cvv"`
}

func (c *CreateTransactionModel) ApplyDiscount() {
	switch c.PaymentMethod {
	case CreditCard:
		c.Value = c.Value - (c.Value * 0.05)
	case DebitCard:
		c.Value = c.Value - (c.Value * 0.03)
	}
}

type PaymentMethod uint8

const (
	CreditCard PaymentMethod = iota
	DebitCard
)
