package money

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
	return m.Amount() == mm.Amount()
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) times(multiplier int) *Money {
	return NewMoney(m.Amount()*multiplier, m.currency)
}

type Dollar struct {
	Money
}

func NewDollar(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "USD",
	}
}

func (d *Dollar) amount() int {
	return d.Money.amount
}

func (d *Dollar) times(multiplier int) *Money {
	return NewMoney(d.amount()*multiplier, d.currency)
}

type Franc struct {
	Money
}

func NewFranc(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "CHF",
	}
}

func (f *Franc) times(multiplier int) *Money {
	return NewMoney(f.amount()*multiplier, f.currency)
}

func (f *Franc) amount() int {
	return f.Money.amount
}
