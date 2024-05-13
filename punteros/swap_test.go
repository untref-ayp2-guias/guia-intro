package punteros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	a, b := 78, 86

	Swap(&a, &b)
	assert.Equal(t, 86, a)
	assert.Equal(t, 78, b)
}
