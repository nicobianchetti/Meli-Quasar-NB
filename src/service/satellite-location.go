package service

import "github.com/nicobianchetti/Meli-Quasar-NB/src/model"

const (
	xKENOBI    = -500.0
	yKENOBI    = -200.0
	xSKYWALKER = 100.0
	ySKYWALKER = -100.0
	xSATO      = 500.0
	ySATO      = 100.0
)

func location(location chan resultLocation, distances []model.DataDistances) {

	defer close(location)

	// time.Sleep(4 * time.Second)

	x, y := getLocation(distances...)

	location <- resultLocation{x: x, y: y}

}

//GetLocation .
// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func getLocation(distances ...model.DataDistances) (x, y float64) {

	var c1, c2, c3 model.Cordinate

	for _, v := range distances {
		switch v.Satellite {
		case "kenobi":
			c1 = model.Cordinate{X: xKENOBI, Y: yKENOBI, D: v.Distance}
		case "skywalker":
			c2 = model.Cordinate{X: xSKYWALKER, Y: ySKYWALKER, D: v.Distance}
		case "sato":
			c3 = model.Cordinate{X: xSATO, Y: ySATO, D: v.Distance}
		}

	}

	return Trilateration(c1, c2, c3)
}
