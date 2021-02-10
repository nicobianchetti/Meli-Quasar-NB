package service

import (
	"errors"
	"reflect"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

type resultLocation struct {
	x, y float64
}

type resultMessage struct {
	message string
	err     error
}

type satelliteService struct {
	cache interfaces.ISatelliteCache
}

//NewSatelliteService .
func NewSatelliteService(cache interfaces.ISatelliteCache) interfaces.ISatelliteService {
	return &satelliteService{cache}
}

func (s *satelliteService) GetTransmitter(satellite *[]model.Satellite) (*model.DTOResult, error) {

	controlSatellite := make(map[string]int) //controla que esten llegando los 3 satélites requeridos

	for _, v := range *satellite {
		controlSatellite[v.Name]++
	}

	sat1 := "kenobi"
	sat2 := "skywalker"
	sat3 := "sato"

	for i, v := range controlSatellite {
		if v != 1 || (i != sat1 && i != sat2 && i != sat3) {
			err := errors.New("Error en satélites ingresados")
			return nil, err
		}
	}

	//------------------------------------------------------------------------------------------

	//Si llegaron datos de los 3 satélites
	var distances []model.DataDistances
	var messages [][]string
	chanLocation := make(chan resultLocation)
	chanMessage := make(chan resultMessage)

	for _, v := range *satellite {
		distances = append(distances, model.DataDistances{Satellite: v.Name, Distance: v.Distance})
		messages = append(messages, v.Message)
	}

	go location(chanLocation, distances)
	go message(chanMessage, messages)

	//Lectura de chanel de Location
	var result resultLocation
	for u := range chanLocation {
		result = resultLocation{x: u.x, y: u.y}
	}

	position := model.NewPosition(result.x, result.y)

	//Lectura chanel de Message
	var err error
	var message string
	for m := range chanMessage {
		message = m.message
		err = m.err
	}

	if err != nil {
		return nil, err
	}

	return model.NewDTOResult(position, message), nil

	// //Sin uso de goroutines
	// x, y := GetLocation(distances...)

	// position := model.NewPosition(x, y)

	// message, err := GetMessage(messages...)
}

func (s *satelliteService) RegisterKey(key string, satellite *model.Satellite) error {

	satellites, err := s.cache.Get(key)

	if err != nil {
		return err
	}

	//Si ya había una estructura de satélite guardada para el usuario, reviso para sobreescribir o agregar satélite nuevo
	if satellites != nil {
		var isExistSatellite bool = false
		var indexExist int
		var updateData bool = false

		//Verifico, si ya existe , sobreescribo los datos y reemplazo la key
		for i, v := range satellites.Satellites {
			if v.Name == satellite.Name {
				isExistSatellite = true
				if !reflect.DeepEqual(v.Message, satellite.Message) || v.Distance != satellite.Distance {
					updateData = true
				}
				indexExist = i
				break
			}
		}

		//Si satélite nuevo no existía , apependeo a estructura
		if !isExistSatellite {
			satellites.Satellites = append(satellites.Satellites, *satellite)
		} else {
			//Si ya existe, y además vino un dato distino,lo modifico con los datos nuevos
			if updateData {
				satellites.Satellites[indexExist].Distance = satellite.Distance
				satellites.Satellites[indexExist].Message = satellite.Message
			}
		}

		err := s.cache.Delete(key)

		if err != nil {
			return err
		}

		err = s.cache.Set(key, satellites)

		if err != nil {
			return err
		}

		return nil
	}

	//Si la estructura no existe , agrego a la base estructura nueva asociada a key(user)
	newSatellites := model.DTORequestSatellites{}

	newSatellites.Satellites = append(newSatellites.Satellites, *satellite)

	err = s.cache.Set(key, &newSatellites)

	if err != nil {
		return err
	}

	return nil

}

func (s *satelliteService) GetSatellites(key string) (*model.DTORequestSatellites, error) {

	satellites, err := s.cache.Get(key)

	if err != nil {
		return nil, err
	}

	if satellites == nil {
		err := errors.New("No hay información suficiente")
		return nil, err
	}

	return satellites, nil
}

func (s *satelliteService) DeleteKey(key string) error {

	err := s.cache.Delete(key)

	if err != nil {
		return nil
	}

	return nil
}
