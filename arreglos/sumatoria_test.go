package arreglos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumatoria(t *testing.T) {
	tests := []struct {
		want int
		args []int
	}{
		{
			want: 0,
			args: []int{},
		},
		{
			want: 42,
			args: []int{42},
		},
		{
			want: 23,
			args: []int{4, 7, 2, 9, 1},
		},
		{
			want: 17,
			args: []int{4, 7, -2, 9, -1},
		},
		{
			want: -17,
			args: []int{-4, -7, 2, -9, 1},
		},
	}

	for _, test := range tests {
		got := Sumatoria(test.args)
		assert.Equal(t, test.want, got)
	}
}
