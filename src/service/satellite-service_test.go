package service

import (
	"reflect"
	"testing"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRedisCache struct {
	mock.Mock
}

func (m *MockRedisCache) Set(key string, value *model.DTORequestSatellites) error {
	//Stub para devolver argumentos
	args := m.Called()

	//Devolución del primer argumento que recibe

	return args.Error(1)
}

func (m *MockRedisCache) Get(key string) (*model.DTORequestSatellites, error) {
	//Stub para devolver argumentos
	args := m.Called()
	result := args.Get(0)

	//Devolución del primer argumento que recibe

	return result.(*model.DTORequestSatellites), args.Error(1)
}

func (m *MockRedisCache) Delete(key string) error {
	//Stub para devolver argumentos
	args := m.Called()

	//Devolución del primer argumento que recibe

	return args.Error(1)
}

func TestGetTransmitter(t *testing.T) {
	cases := []struct {
		name      string
		satellite []model.Satellite
		result    *model.DTOResult
		err       error
	}{
		{
			name:      "Case 1: Ejemplo enunciado",
			satellite: []model.Satellite{{Name: "kenobi", Distance: 100, Message: []string{"este", "", "", "mensaje", ""}}, {Name: "skywalker", Distance: 115.5, Message: []string{"", "es", "", "", "secreto"}}, {Name: "sato", Distance: 142.7, Message: []string{"este", "", "un", "", ""}}},
			result:    &model.DTOResult{model.Position{X: -487.2859125, Y: 1557.014225}, "este es un mensaje secreto"},
			err:       nil,
		},

		{
			name:      "Case 2: Satélite duplicado ",
			satellite: []model.Satellite{{Name: "kenobi", Distance: 100, Message: []string{"este", "", "", "mensaje", ""}}, {Name: "kenobi", Distance: 115.5, Message: []string{"", "es", "", "", "secreto"}}, {Name: "sato", Distance: 142.7, Message: []string{"este", "", "un", "", ""}}},
			result:    nil,
			err:       errorSatelliteInpunt,
		},

		{
			name:      "Case 3: Satélite incorrecto ",
			satellite: []model.Satellite{{Name: "kenobiHHH", Distance: 100, Message: []string{"este", "", "", "mensaje", ""}}, {Name: "kenobi", Distance: 115.5, Message: []string{"", "es", "", "", "secreto"}}, {Name: "sato", Distance: 142.7, Message: []string{"este", "", "un", "", ""}}},
			result:    nil,
			err:       errorSatelliteInpunt,
		},

		{
			name:      "Case 4: Falta un satélite ",
			satellite: []model.Satellite{{Name: "kenobi", Distance: 100, Message: []string{"este", "", "", "mensaje", ""}}, {Name: "kenobi", Distance: 115.5, Message: []string{"", "es", "", "", "secreto"}}},
			result:    nil,
			err:       errorSatelliteInpunt,
		},

		{
			name:      "Case 5: Una distancia faltante ",
			satellite: []model.Satellite{{Name: "kenobi", Distance: 100, Message: []string{"este", "", "", "mensaje", ""}}, {Name: "kenobi", Distance: 0, Message: []string{"", "es", "", "", "secreto"}}},
			result:    nil,
			err:       errorSatelliteInpunt,
		},
	}

	mockRepo := new(MockRedisCache)

	testService := NewSatelliteService(mockRepo)

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			result, err := testService.GetTransmitter(&test.satellite)

			if !reflect.DeepEqual(result, test.result) {
				t.Error("Retorna:", result, ",y el valor correcto es:", test.result)
			}

			if !reflect.DeepEqual(err, test.err) {
				t.Error("Retorna error:", err, ",y el valor correcto es:", test.err)
			}
		})
	}

}
