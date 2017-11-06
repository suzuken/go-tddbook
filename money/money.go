package money

type IMoney interface {
	amount() int
}

type Money struct {
	amount int
}

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Money{amount}}
}

func (d *Dollar) amount() int {
	return d.Money.amount
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount() * multiplier)
}

func (d *Dollar) equals(object interface{}) bool {
	m := object.(IMoney)
	return d.amount() == m.amount()
}

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{Money{amount}}
}

func (f *Franc) times(multiplier int) *Franc {
	return NewFranc(f.amount() * multiplier)
}

func (f *Franc) amount() int {
	return f.Money.amount
}

func (f *Franc) equals(object interface{}) bool {
	m := object.(IMoney)
	return f.amount() == m.amount()
}
