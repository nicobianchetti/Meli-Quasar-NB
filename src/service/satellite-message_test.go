package service

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMessages(t *testing.T) {
	cases := []struct {
		name     string
		messages [][]string
		result   string
		err      error
	}{
		{
			name:     "Case 1: Ejemplo enunciado",
			messages: [][]string{{"este", "", "", "mensaje", ""}, {"", "es", "", "", "secreto"}, {"este", "", "un", "", ""}},
			result:   "este es un mensaje secreto",
			err:      nil,
		},

		{
			name:     "Case 2: Información incompleta",
			messages: [][]string{{"este", "", "", "mensaje", ""}, {"", "es", "", "", "secreto"}, {"este", "", "", "", ""}},
			result:   "",
			err:      errors.New("Error mensaje. Falta información"),
		},

		{
			name:     "Case 3: Superposición de cadenas",
			messages: [][]string{{"este", "", "", "mensaje", ""}, {"", "es", "uno", "", "secreto"}, {"este", "", "un", "", ""}},
			result:   "",
			err:      errors.New("Error mensaje. Superposición de cadenas"),
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			result, err := GetMessage(test.messages...)

			if result != test.result {
				t.Error("Retorna:", result, ",y el valor correcto es:", test.result)
			}

			fmt.Println(err)
			fmt.Println(test.err)

			if !reflect.DeepEqual(err, test.err) {
				t.Error("Retorna error:", err, ",y el valor correcto es:", test.err)
			}
		})
	}
}
