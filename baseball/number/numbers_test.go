package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNumbers(t *testing.T) {
	numbers, err := NewNumbers("123")
	assert.NotNil(t, numbers)
	assert.Nil(t, err)
}

func TestContains(t *testing.T) {
	numbers, _ := NewNumbers("123")
	assert.Equal(t, Strike, numbers.Contains("1", 0))
	assert.NotEqual(t, Strike, numbers.Contains("1", 2))
}
