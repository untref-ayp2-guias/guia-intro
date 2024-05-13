package bucles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		want int
		args int
	}{
		{
			want: 1,
			args: 0,
		},
		{
			want: 1,
			args: 1,
		},
		{
			want: 120,
			args: 5,
		},
	}

	for _, test := range tests {
		got := Factorial(test.args)
		assert.Equal(t, test.want, got)
	}
}
