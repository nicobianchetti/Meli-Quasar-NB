package service

import (
	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

type satelliteService struct{}

//NewSatelliteService .
func NewSatelliteService() interfaces.ISatelliteService {
	return &satelliteService{}
}

func (s *satelliteService) TopSecret(*[]model.Satellite) *model.DTOResult {
	return nil
}
