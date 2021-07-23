package baseball

import (
	score2 "baseball/baseball/score"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateGame(t *testing.T) {
	game, _ := NewGame("123")
	assert.NotNil(t, game)
}

func TestInvalidNumbers(t *testing.T) {
	assertError(t, "#12")
	assertError(t, "1#2")
	assertError(t, "120")
}

func TestInvalidLength(t *testing.T) {
	assertError(t, "12")
	assertError(t, "1234")
}

func TestDuplicated(t *testing.T) {
	assertError(t, "112")
	assertError(t, "122")
}

func assertError(t *testing.T, s string) {
	game, err := NewGame(s)
	assert.Nil(t, game)
	assert.Error(t, err)
}

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
	assert.Equal(t, score2.New(strike, ball), result)
}
