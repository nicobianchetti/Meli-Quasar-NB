package service

import (
	"testing"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

func TestTrilateration(t *testing.T) {
	cases := []struct {
		name       string
		c1, c2, c3 model.Cordinate
		x, y       float64
	}{
		{
			name: "Case 1: Ejemplo enunciado",
			c1:   model.Cordinate{X: -500, Y: -200, D: 100},
			c2:   model.Cordinate{X: 100, Y: -100, D: 115.5},
			c3:   model.Cordinate{X: 500, Y: 100, D: 142.7},
			x:    -487.2859125,
			y:    1557.014225,
		},
		{
			name: "Case 2: 3 satélites - 2 distancias",
			c1:   model.Cordinate{X: -500, Y: -200, D: 100},
			c2:   model.Cordinate{X: 100, Y: -100, D: 115.5},
			c3:   model.Cordinate{X: 500, Y: 100, D: 0},
			x:    -500.01296875,
			y:    1633.3765625,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			x, y := Trilateration(test.c1, test.c2, test.c3)

			if x != test.x {
				t.Error("Retorna valor en x:", x, ",y el valor correcto es:", test.x)
			}

			if y != test.y {
				t.Error("Retorna valor en x:", y, ",y el valor correcto es:", test.y)
			}
		})
	}
}
