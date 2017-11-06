package money

import "fmt"

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]int),
	}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(*b, to)
}

func (b *Bank) addRate(from, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

func (b *Bank) rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[Pair{from, to}]
}

type Pair struct {
	from, to string
}

type Expression interface {
	Reduce(bank Bank, to string) *Money
}

type Sum struct {
	addend, augend Expression
}

func (s *Sum) Reduce(bank Bank, to string) *Money {
	amount := s.augend.Reduce(bank, to).Amount() +
		s.addend.Reduce(bank, to).Amount()
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

func (m *Money) times(multiplier int) Expression {
	return NewMoney(m.Amount()*multiplier, m.currency)
}

func (m *Money) plus(addend Expression) Expression {
	return &Sum{m, addend}
}

func (m *Money) Reduce(bank Bank, to string) *Money {
	rate := bank.rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
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
