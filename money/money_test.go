package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.amount != 10 {
		t.Errorf("want 10, got %d", product.amount)
	}
	product = five.times(3)
	if product.amount != 15 {
		t.Errorf("want 15, got %d", product.amount)
	}
}
