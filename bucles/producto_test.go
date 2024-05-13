package bucles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProducto(t *testing.T) {
	tests := []struct {
		want int
		args [2]int
	}{
		{
			want: 12,
			args: [2]int{3, 4},
		},
		{
			want: -12,
			args: [2]int{3, -4},
		},
		{
			want: -12,
			args: [2]int{-3, 4},
		},
		{
			want: 12,
			args: [2]int{-3, -4},
		},
		{
			want: 0,
			args: [2]int{-3, 0},
		},
		{
			want: 0,
			args: [2]int{0, 0},
		},
	}

	for _, test := range tests {
		got := Producto(test.args[0], test.args[1])
		assert.Equal(t, test.want, got, "Producto(%d, %d)", test.args[0], test.args[1])
	}
}
