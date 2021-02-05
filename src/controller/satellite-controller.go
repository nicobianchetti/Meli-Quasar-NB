package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

type satelliteController struct{}

//NewSatelliteController .
func NewSatelliteController() interfaces.ISatelliteController {
	return &satelliteController{}
}

func (s *satelliteController) TopSecret(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var satellitesInput model.DTORequestSatellites
	err := decoder.Decode(&satellitesInput)

	defer r.Body.Close()

	if err != nil {
		fmt.Println("\n error decoding:", err)
	}

	fmt.Println(satellitesInput)
}
