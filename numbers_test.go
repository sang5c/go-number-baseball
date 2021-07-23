package main

import (
	"baseball/score"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	game, _ = NewGame("123")
)

func TestNothing(t *testing.T) {
	assertEqual(t, "456", 0, 0)
}

func TestStrike(t *testing.T) {
	assertEqual(t, "123", 3, 0)
	assertEqual(t, "143", 2, 0)
	assertEqual(t, "425", 1, 0)
}

func assertEqual(t *testing.T, input string, strike int, ball int) {
	result := game.compare(input)
	assert.Equal(t, score.New(strike, ball), result)
}
