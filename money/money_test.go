package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	assert.Equal(t, product.amount, 10)
	product = five.times(3)
	assert.Equal(t, product.amount, 15)
}

func TestMoneyEquality(t *testing.T) {
	d := Dollar{5}
	assert.Equal(t, d.equals(Dollar{5}), true)
	assert.Equal(t, d.equals(Dollar{6}), false)
}
