package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.times(2))
	assert.Equal(t, NewDollar(15), five.times(3))
}

func TestMoneyEquality(t *testing.T) {
	assert.Equal(t, true, NewDollar(5).equals(NewDollar(5)))
	assert.Equal(t, false, NewDollar(5).equals(NewDollar(6)))
	assert.Equal(t, true, NewFranc(5).equals(NewFranc(5)))
	assert.Equal(t, false, NewFranc(5).equals(NewFranc(6)))
	assert.Equal(t, false, NewFranc(5).equals(NewDollar(5)))
}

func TestMoneyFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.times(2))
	assert.Equal(t, NewFranc(15), five.times(3))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

// Oh it's not class ;)
func TestDifferentClassEquality(t *testing.T) {
	assert.Equal(t, NewMoney(10, "CHF").equals(NewMoney(10, "CHF")), true)
}
