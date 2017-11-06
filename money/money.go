package money

import "fmt"

type Bank struct{}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(to)
}

type Expression interface {
	Reduce(to string) *Money
}

type Sum struct {
	addend, augend *Money
}

func (s *Sum) Reduce(to string) *Money {
	amount := s.augend.amount + s.addend.amount
	return NewMoney(amount, to)
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
}

func (m *Money) Reduce(to string) *Money {
	return m
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
