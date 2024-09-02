package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, float64(5), Max(1, 5))
	assert.Equal(t, float64(-40), Max(-100, -40))
	assert.Equal(t, 5.5, Max(5.5, 5.1))
}
