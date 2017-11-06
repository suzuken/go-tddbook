package money

type IMoney interface {
	Amount() int
	Currency() string
}

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int) *Money {
	return &Money{
		amount: amount,
	}
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) equals(object interface{}) bool {
	mm := object.(IMoney)
	return m.Amount() == mm.Amount()
}

func (m *Money) times(multiplier int) *Money {
	return NewMoney(m.Amount() * multiplier)
}

type Dollar struct {
	Money
}

func NewDollar(amount int, currency string) *Dollar {
	return &Dollar{Money{
		amount:   amount,
		currency: "USD",
	}}
}

func (d *Dollar) amount() int {
	return d.Money.amount
}

func (d *Dollar) Currency() string {
	return d.Money.currency
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount()*multiplier, "")
}

type Franc struct {
	Money
}

func NewFranc(amount int, currency string) *Franc {
	return &Franc{Money{
		amount:   amount,
		currency: "CHF",
	}}
}

func (f *Franc) Currency() string {
	return f.Money.currency
}

func (f *Franc) times(multiplier int) *Franc {
	return NewFranc(f.amount()*multiplier, "")
}

func (f *Franc) amount() int {
	return f.Money.amount
}
