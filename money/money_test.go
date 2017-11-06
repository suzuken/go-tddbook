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
	assert.Equal(t, NewDollar(5).equals(NewDollar(5)), true)
	assert.Equal(t, NewDollar(5).equals(NewDollar(6)), false)
	assert.Equal(t, NewFranc(5).equals(NewFranc(5)), true)
	assert.Equal(t, NewFranc(5).equals(NewFranc(6)), false)
}

func TestMoneyFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, NewFranc(10), five.times(2))
	assert.Equal(t, NewFranc(15), five.times(3))
}
