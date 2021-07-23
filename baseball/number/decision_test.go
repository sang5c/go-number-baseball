package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecision(t *testing.T) {
	assert.True(t, Strike.IsStrike())
	assert.False(t, Strike.IsBall())
	assert.True(t, Ball.IsBall())
	assert.False(t, Ball.IsStrike())
	assert.False(t, Nothing.IsStrike())
	assert.False(t, Nothing.IsBall())
}
