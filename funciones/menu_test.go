package funciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarOpcion(t *testing.T) {
	tests := []struct {
		want string
		args string
	}{
		{
			want: "Se ha seleccionado la opción A",
			args: "A",
		},
		{
			want: "Se ha seleccionado la opción C",
			args: "C",
		},
		{
			want: "La opción X es incorrecta",
			args: "X",
		},
	}

	for _, test := range tests {
		got := ValidarOpcion(test.args)
		assert.Equal(t, test.want, got)
	}
}
