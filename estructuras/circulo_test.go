package estructuras

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCirculo(t *testing.T) {
	centro := NewPunto(0, 0)
	circulo := NewCirculo(centro, 3.0)

	assert.InDelta(t, 28.25, circulo.Area(), 0.05)
	assert.InDelta(t, 18.84, circulo.Perimetro(), 0.05)
	assert.Equal(t, "Círculo: {centro: (0, 0), radio: 3}", circulo.String())
}
