package estructuras

import "fmt"

type Punto struct {
	X float64
	Y float64
}

// NewPunto devuelve un nuevo punto con coordenadas (x, y).
func NewPunto(x, y float64) *Punto {
	return &Punto{X: x, Y: y}
}

// String devuelve la representación en cadena del Punto.
func (p *Punto) String() string {
	return fmt.Sprintf("Punto: (%v, %v)", p.X, p.Y)
}
