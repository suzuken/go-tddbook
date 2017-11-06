package money

import "fmt"

type Bank struct{}

func (b *Bank) reduce(source Expression, to string) *Money {
	sum, ok := source.(*Sum)
	if !ok {
		return nil
	}
	amount := sum.augend.amount + sum.addend.amount
	return NewMoney(amount, to)
}

type Expression interface{}

type Sum struct {
	addend, augend *Money
}

type IMoney interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) equals(object interface{}) bool {
	mm := object.(IMoney)
	return m.Amount() == mm.Amount() &&
		m.Currency() == mm.Currency()
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) times(multiplier int) *Money {
	return NewMoney(m.Amount()*multiplier, m.currency)
}

func (m *Money) plus(addend *Money) Expression {
	return &Sum{m, addend}
	// return NewMoney(m.amount+addend.amount, m.currency)
}

func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}

func NewDollar(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "USD",
	}
}

func NewFranc(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "CHF",
	}
}
