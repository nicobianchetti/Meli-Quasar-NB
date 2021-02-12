package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

type satelliteController struct {
	service interfaces.ISatelliteService
}

var (
	errorNombreSatellite  = errors.New("Error. Nombre de satélite incorrecto")
	errorUsuario          = errors.New("Error. Ingrese usuario")
	errorInfoInsuficiente = errors.New("No hay información suficiente")
)

//NewSatelliteController .
func NewSatelliteController(service interfaces.ISatelliteService) interfaces.ISatelliteController {
	return &satelliteController{service}
}

func (s *satelliteController) TopSecret(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var satellitesInput model.DTORequestSatellites
	err := decoder.Decode(&satellitesInput)

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

func (s *satelliteController) TopSecretSplit(w http.ResponseWriter, r *http.Request) {
	nameSatellite := mux.Vars(r)["id"]

	if nameSatellite != "kenobi" && nameSatellite != "skywalker" && nameSatellite != "sato" {
		http.Error(w, errorNombreSatellite.Error(), http.StatusNotFound)
		return
	}

	user := r.Header.Get("user")

	if user == "" {
		http.Error(w, errorUsuario.Error(), http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var satelliteInput model.Satellite
	err := decoder.Decode(&satelliteInput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if satelliteInput.Distance == 0 {
		err := errors.New("Error. Ingrese distancia")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	satelliteInput.Name = nameSatellite

	err = s.service.RegisterKey(user, &satelliteInput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, nil)

}

func (s *satelliteController) TopSecretSplitGet(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("user")

	if user == "" {
		http.Error(w, errorUsuario.Error(), http.StatusNotFound)
		return
	}

	satellites, err := s.service.GetSatellites(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if satellites == nil {
		http.Error(w, errorInfoInsuficiente.Error(), http.StatusNotFound)
		return
	}

	if len(satellites.Satellites) != 3 {
		http.Error(w, errorInfoInsuficiente.Error(), http.StatusNotFound)
		return
	}

	result, err := s.service.GetTransmitter(&satellites.Satellites)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	//Limpio set de datos ingresados para ese usuario
	err = s.service.DeleteKey(user)

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

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
