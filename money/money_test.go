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
	d := Dollar{5}
	assert.Equal(t, d.equals(Dollar{5}), true)
	assert.Equal(t, d.equals(Dollar{6}), false)
}

func TestMoneyFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.times(2))
	assert.Equal(t, NewFranc(15), five.times(3))
}
