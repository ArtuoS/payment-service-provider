package domain

import (
	"strings"
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

type GetTransactionResult struct {
	ID                 int64         `db:"id" json:"id"`
	Value              float32       `db:"transaction_value" json:"transaction_value"`
	Description        string        `db:"description" json:"description"`
	PaymentMethodDb    PaymentMethod `db:"payment_method" json:"-"`
	PaymentMethod      string        `json:"payment_method"`
	CardNumberDb       string        `db:"card_number" json:"-"`
	CardNumber         string        `json:"card_number"`
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

func (g *GetTransactionResult) Format() {
	g.CardNumber = strings.Repeat("*", len(g.CardNumberDb)-4) + g.CardNumberDb[len(g.CardNumberDb)-4:]
	g.PaymentMethod = g.PaymentMethodDb.String()
}

type PaymentMethod uint8

const (
	CreditCard PaymentMethod = iota
	DebitCard
)

func (p PaymentMethod) String() string {
	switch p {
	case CreditCard:
		return "credit_card"
	case DebitCard:
		return "debit_card"
	}
	return ""
}
