package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDec2Bin(t *testing.T) {
	tests := []struct {
		want string
		args int
	}{
		{
			want: "0",
			args: 0,
		},
		{
			want: "1",
			args: 1,
		},
		{
			want: "1010",
			args: 10,
		},
		{
			want: "101010",
			args: 42,
		},
		{
			want: "10100111001",
			args: 1337,
		},
	}

	for _, test := range tests {
		got := Dec2Bin(test.args)
		assert.Equal(t, test.want, got)
	}
}
