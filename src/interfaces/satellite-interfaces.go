package interfaces

import (
	"net/http"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

//ISatelliteController .
type ISatelliteController interface {
	TopSecret(w http.ResponseWriter, r *http.Request)
}

//ISatelliteService .
type ISatelliteService interface {
	TopSecret(*[]model.Satellite) *model.DTOResult
}

// //ISatelliteRepository .
// type ISatelliteRepository interface {
// }
