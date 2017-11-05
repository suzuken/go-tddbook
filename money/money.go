package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) equals(object interface{}) bool {
	dollar := object.(Dollar)
	return d.amount == dollar.amount
}

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{
		amount: amount,
	}
}

func (d *Franc) times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}

func (d *Franc) equals(object interface{}) bool {
	franc := object.(Franc)
	return d.amount == franc.amount
}
