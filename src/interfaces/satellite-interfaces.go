package interfaces

import (
	"net/http"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

//ISatelliteController .
type ISatelliteController interface {
	TopSecret(w http.ResponseWriter, r *http.Request)
	TopSecretSplit(w http.ResponseWriter, r *http.Request)
	TopSecretSplitGet(w http.ResponseWriter, r *http.Request)
}

//ISatelliteService .
type ISatelliteService interface {
	GetTransmitter(*[]model.Satellite) (*model.DTOResult, error)
	RegisterKey(string, *model.Satellite) error
	GetSatellites(string) (*model.DTORequestSatellites, error)
	DeleteKey(string) error
}

//ISatelliteCache .
type ISatelliteCache interface {
	Set(key string, value *model.DTORequestSatellites) error
	Get(key string) (*model.DTORequestSatellites, error)
	Delete(key string) error
}
