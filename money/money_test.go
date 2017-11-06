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
	assert.Equal(t, false, NewFranc(5).equals(NewDollar(5)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.plus(five)
	bank := Bank{}
	reduced := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturns(t *testing.T) {
	five := NewDollar(5)
	result := five.plus(five)
	sum, ok := result.(*Sum)
	if !ok {
		t.Fatal("want Sum")
	}
	assert.Equal(t, five, sum.augend)
	assert.Equal(t, five, sum.addend)
}

func TestReduceSum(t *testing.T) {
	sum := &Sum{NewDollar(3), NewDollar(4)}
	bank := Bank{}
	result := bank.reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := Bank{}
	result := bank.reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}
