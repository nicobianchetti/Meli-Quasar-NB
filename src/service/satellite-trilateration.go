package service

import (
	"math"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

//Trilateration .
func Trilateration(c1, c2, c3 model.Cordinate) (x, y float64) {
	// https://confluence.slac.stanford.edu/display/IEPM/TULIP+Algorithm+Alternative+Trilateration+Method

	// fmt.Println("\n c1:", c1)
	// fmt.Println("\n c2:", c2)
	// fmt.Println("\n c3:", c3)

	d1, d2, d3, i1, i2, i3, j1, j2, j3 := &c1.D, &c2.D, &c3.D, &c1.X, &c2.X, &c3.X, &c1.Y, &c2.Y, &c3.Y

	denom1 := ((2**i2-2**i3)*(2**j2-2**j1) - (2**i1-2**i2)*(2**j3-2**j2))
	denom2 := (2**j2 - 2**j1)

	//Me aseguro de no generar ninguna división por 0
	if denom1 != 0 && denom2 != 0 {
		x = ((((math.Pow(*d1, 2)-math.Pow(*d2, 2))+(math.Pow(*i2, 2)-math.Pow(*i1, 2))+(math.Pow(*j2, 2)-math.Pow(*j1, 2)))*(2**j3-2**j2) - ((math.Pow(*d2, 2)-math.Pow(*d3, 2))+(math.Pow(*i3, 2)-math.Pow(*i2, 2))+(math.Pow(*j3, 2)-math.Pow(*j2, 2)))*(2**j2-2**j1)) / denom1)

		y = ((math.Pow(*d1, 2) - math.Pow(*d2, 2)) + (math.Pow(*i2, 2) - math.Pow(*i1, 2)) + (math.Pow(*j2, 2) - math.Pow(*j1, 2)) + x*(2**i1-2**i2)) / denom2

		// fmt.Println("aca x:", x, " y:", y)

		return x, y
	}

	return 0, 0
}
