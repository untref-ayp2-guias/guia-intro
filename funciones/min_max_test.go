package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		wantMin int
		wantMax int
		args    []int
	}{
		{
			wantMin: 1,
			wantMax: 9,
			args:    []int{4, 7, 2, 9, 1},
		},
		{
			wantMin: -9,
			wantMax: -1,
			args:    []int{-4, -7, -2, -9, -1},
		},
		{
			wantMin: 4,
			wantMax: 4,
			args:    []int{4},
		},
		{
			wantMin: 5,
			wantMax: 5,
			args:    []int{5, 5, 5, 5, 5},
		},
	}

	for _, test := range tests {
		gotMin, gotMax := MinMax(test.args)
		assert.Equal(t, test.wantMin, gotMin)
		assert.Equal(t, test.wantMax, gotMax)
	}
}
