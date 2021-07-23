package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	number, err := NewNumber("1", 0)
	assert.NotNil(t, number)
	assert.Nil(t, err)
}

func TestCompareBall(t *testing.T) {
	source, _ := NewNumber("1", 0)
	target, _ := NewNumber("1", 1)
	result := source.compare(target)

	assert.Equal(t, Ball, result)
}

func TestCompareStrike(t *testing.T) {
	source, _ := NewNumber("1", 0)
	target, _ := NewNumber("1", 0)
	result := source.compare(target)

	assert.Equal(t, Strike, result)
}
