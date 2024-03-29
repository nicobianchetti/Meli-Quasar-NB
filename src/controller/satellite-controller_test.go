package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockSatelliteService struct {
	mock.Mock
}

func (m *mockSatelliteService) GetTransmitter(*[]model.Satellite) (*model.DTOResult, error) {
	//Stub para devolver argumentos
	args := m.Called()

	//Devolución del primer argumento que recibe
	result := args.Get(0)

	return result.(*model.DTOResult), args.Error(1)
}

func (m *mockSatelliteService) RegisterKey(string, *model.Satellite) error {
	//Stub para devolver argumentos
	args := m.Called()

	//Devolución del primer argumento que recibe

	return args.Error(1)
}

func (m *mockSatelliteService) GetSatellites(string) (*model.DTORequestSatellites, error) {
	return nil, nil
}

func (m *mockSatelliteService) DeleteKey(string) error {
	return nil
}

func TestTopSecret(t *testing.T) {
	//Mockeo Service para pasarlo por inyección de dependencia al controller
	mockService := new(mockSatelliteService)

	position := model.Position{
		X: -487.2859125,
		Y: 1557.014225,
	}
	satellite := model.DTOResult{
		Position: position,
		Message:  "este es un mensaje secreto",
	}

	mockService.On("GetTransmitter").Return(&satellite, nil)

	testController := NewSatelliteController(mockService)

	//----------------------------------------------------------

	var jsonReq = []byte(`{"satellites":[{"name": "kenobi","distance": 100.00,"message": ["este", "", "", "mensaje", ""]},{"name": "skywalker","distance": 115.5 ,"message": ["", "es", "", "", "secreto"]},{"name": "sato","distance": 142.7 ,"message": ["este", "", "un", "", ""]}]}`)

	req, err := http.NewRequest(http.MethodPost, "quasar/topsecret/", bytes.NewBuffer(jsonReq))

	if err != nil {
		t.Error("Error in request test: ", err.Error())
	}

	//Asing HTTP Handler function(controller Create function)
	handler := http.HandlerFunc(testController.TopSecret)

	//Record HTTP Response (with httptest library)
	response := httptest.NewRecorder()

	//Dispach the HTTP Request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code ant the response
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	var satelliteRes *model.DTOResult

	json.NewDecoder(io.Reader(response.Body)).Decode(&satelliteRes)

	//Assert HTTP response
	assert.NotNil(t, satelliteRes.Message)
	assert.NotNil(t, satelliteRes.Position.X)
	assert.NotNil(t, satelliteRes.Position.Y)
	assert.Equal(t, (*satelliteRes).Message, "este es un mensaje secreto")
	assert.Equal(t, (*&satelliteRes).Position.X, -487.2859125)
	assert.Equal(t, (*&satelliteRes).Position.Y, 1557.014225)

}
