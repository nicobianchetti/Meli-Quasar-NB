package controller

import (
	"encoding/json"
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
