package main

import (
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

func assertError(t *testing.T, s string) {
	game, err := NewGame(s)
	assert.Nil(t, game)
	assert.Error(t, err)
}
