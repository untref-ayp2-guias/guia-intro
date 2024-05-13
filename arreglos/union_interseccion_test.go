package arreglos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionInterseccion(t *testing.T) {
	tests := []struct {
		wantUnion        []int
		wantInterseccion []int
		args             [2][]int
	}{
		{
			wantUnion:        []int{},
			wantInterseccion: []int{},
			args:             [2][]int{{}, {}},
		},
		{
			wantUnion:        []int{3, 4},
			wantInterseccion: []int{3},
			args:             [2][]int{{3, 3}, {3, 4, 4}},
		},
		{
			wantUnion:        []int{1, 2, 3, 4, 5, 6, 7},
			wantInterseccion: []int{3, 4, 5},
			args:             [2][]int{{1, 2, 3, 4, 5}, {3, 4, 5, 6, 7}},
		},
		{
			wantUnion:        []int{4, 5, 6, 7},
			wantInterseccion: []int{},
			args:             [2][]int{{}, {4, 5, 6, 7}},
		},
		{
			wantUnion:        []int{1, 2, 3, 4, 5, 6},
			wantInterseccion: []int{},
			args:             [2][]int{{1, 2, 3}, {4, 5, 6}},
		},
	}

	for _, test := range tests {
		gotunion, gotInterseccion := UnionInterseccion(test.args[0], test.args[1])
		assert.Equal(t, test.wantUnion, gotunion)
		assert.Equal(t, test.wantInterseccion, gotInterseccion)
	}
}
