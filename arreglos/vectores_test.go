package arreglos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumaVectorialProductoEscalar(t *testing.T) {
	tests := []struct {
		wantSuma     []int
		wantProducto int
		args         [2][]int
	}{
		{
			wantSuma:     []int{},
			wantProducto: 0,
			args:         [2][]int{{}, {}},
		},
		{
			wantSuma:     []int{5, 7, 9},
			wantProducto: 32,
			args:         [2][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			wantSuma:     []int{3, 7, 3},
			wantProducto: -12,
			args:         [2][]int{{-1, 2, -3}, {4, 5, 6}},
		},
	}

	for _, test := range tests {
		gotSuma, gotProducto := SumaVectorialProductoEscalar(test.args[0], test.args[1])
		assert.Equal(t, test.wantSuma, gotSuma)
		assert.Equal(t, test.wantProducto, gotProducto)
	}
}
