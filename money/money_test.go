package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, NewDollar(10), five.Times(2))
	assert.Equal(t, NewDollar(15), five.Times(3))
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
	sum := five.Plus(five)
	bank := Bank{}
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturns(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum, ok := result.(*Sum)
	if !ok {
		t.Fatal("want Sum")
	}
	assert.Equal(t, five, sum.augend)
	assert.Equal(t, five, sum.addend)
}

func TestReduceSum(t *testing.T) {
	sum := &Sum{NewDollar(3), NewDollar(4)}
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestIdentifyRAte(t *testing.T) {
	bank := NewBank()
	assert.Equal(t, 1, bank.rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, NewDollar(10), result)
}

func TestSumPlusMoney(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	// $5 + 10CHF + $5
	sum := &Sum{fiveBucks, tenFrancs}
	exp := sum.Plus(fiveBucks)
	result := bank.Reduce(exp, "USD")
	assert.Equal(t, NewDollar(15), result)
}

func TestSumTimes(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.addRate("CHF", "USD", 2)
	sum := &Sum{fiveBucks, tenFrancs}
	exp := sum.Times(2)
	result := bank.Reduce(exp, "USD")
	assert.Equal(t, NewDollar(20), result)
}
