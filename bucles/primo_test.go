package bucles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEsPrimo(t *testing.T) {
	tests := []struct {
		want bool
		args int
	}{
		{
			want: false,
			args: -3,
		},
		{
			want: false,
			args: 0,
		},
		{
			want: false,
			args: 1,
		},
		{
			want: false,
			args: 12,
		},
		{
			want: false,
			args: 1337,
		},
		{
			want: true,
			args: 5,
		},
		{
			want: true,
			args: 9973,
		},
	}

	for _, test := range tests {
		got := EsPrimo(test.args)
		assert.Equalf(t, test.want, got, "EsPrimo(%d)", test.args)
	}
}
