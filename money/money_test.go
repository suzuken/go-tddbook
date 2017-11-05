package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar{5}
	five.times(2)
	if five.amount != 10 {
		t.Errorf("want 10, got %d", five.amount)
	}
}
