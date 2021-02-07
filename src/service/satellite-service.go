package service

import (
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

type satelliteService struct{}

//NewSatelliteService .
func NewSatelliteService() interfaces.ISatelliteService {
	return &satelliteService{}
}

func (s *satelliteService) GetTransmitter(satellite *[]model.Satellite) (*model.DTOResult, error) {

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

func location(location chan resultLocation, distances []model.DataDistances) {

	defer close(location)

	// time.Sleep(4 * time.Second)

	x, y := GetLocation(distances...)

	location <- resultLocation{x: x, y: y}

}

func message(message chan resultMessage, messages [][]string) {

	defer close(message)

	messageResult, err := GetMessage(messages...)

	message <- resultMessage{message: messageResult, err: err}

}
