package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepresentarPolinomio(t *testing.T) {
	tests := []struct {
		want string
		args []float32
	}{
		{
			want: "3.5+2.5x+1.5x^2",
			args: []float32{3.5, 2.5, 1.5},
		},
		{
			want: "3.5+2.5x-1.5x^2",
			args: []float32{3.5, 2.5, -1.5},
		},
		{
			want: "1+1.5x^2",
			args: []float32{1, 0, 1.5},
		},
		{
			want: "-1+1.5x^2",
			args: []float32{-1, 0, 1.5},
		},
		{
			want: "-3x^3",
			args: []float32{0, 0, 0, -3.0},
		},
		{
			want: "5x^3",
			args: []float32{0, 0, 0, 5.0},
		},
		{
			want: "",
			args: []float32{},
		},
	}

	for _, test := range tests {
		got := RepresentarPolinomio(test.args)
		assert.Equal(t, test.want, got)
	}
}
