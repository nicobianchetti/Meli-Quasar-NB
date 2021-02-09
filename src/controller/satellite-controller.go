package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

type satelliteController struct {
	service interfaces.ISatelliteService
}

//NewSatelliteController .
func NewSatelliteController(service interfaces.ISatelliteService) interfaces.ISatelliteController {
	return &satelliteController{service}
}

func (s *satelliteController) TopSecret(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var satellitesInput model.DTORequestSatellites
	err := decoder.Decode(&satellitesInput)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	controlSatellite := make(map[string]int) //controla que esten llegando los 3 satélites requeridos

	for _, v := range satellitesInput.Satellites {
		controlSatellite[v.Name]++
	}

	sat1 := "kenobi"
	sat2 := "skywalker"
	sat3 := "sato"

	for i, v := range controlSatellite {
		if v != 1 || (i != sat1 && i != sat2 && i != sat3) {
			err := errors.New("Error en los satélites ingresados")
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	result, err := s.service.GetTransmitter(&satellitesInput.Satellites)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
