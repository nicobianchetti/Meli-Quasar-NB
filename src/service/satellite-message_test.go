package service

import (
	"errors"
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

		{
			name:     "Case 4: Desfasaje ejemplo enunciado",
			messages: [][]string{{"", "este", "es", "un", "mensaje"}, {"este", "", "un", "mensaje"}, {"", "", "es", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 5: Desfasaje 2 - Mínimo referencia sin strings ",
			messages: [][]string{{"este", "es", "", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "un", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 5.1: Desfasaje 2 - Mínimo referencia sin strings ",
			messages: [][]string{{"", "", "es", "", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "este", "", "un", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 5.2: Desfasaje 2 - Mínimo sin strings , no reference ",
			messages: [][]string{{"si", "", "es", "", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "este", "", "un", "mensaje"}},
			result:   "si este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 5.3: Desfasaje 2 - Mínimo sin strings , no reference ",
			messages: [][]string{{"", "si", "", "es", "", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "este", "", "un", "mensaje"}},
			result:   "si este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 5.4: Desfasaje 2 - Indeterminación por no tener límites (L/R) ",
			messages: [][]string{{"", "si", "", "es", "", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "este", "", "un", "mensaje", "", ""}},
			result:   "",
			err:      errors.New("Error mensaje. Superposición de cadenas"),
		},

		{
			name:     "Case 6: Desfasaje 3  ",
			messages: [][]string{{"", "", "", "", "es", "un", ""}, {"este", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 6: Desfasaje 4 - Indeterminado  ",
			messages: [][]string{{"", "", "", "", "es", "un"}, {"este", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "",
			err:      errors.New("Error mensaje. Falta información"),
		},

		{
			name:     "Case 7: Desfasaje 5 - Minimo no reference  ",
			messages: [][]string{{"este", "", "", "", "es", "un"}, {"", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "",
			err:      errors.New("Error mensaje. Falta información"),
		},

		{
			name:     "Case 8: Desfasaje 6 - Doble referencia array menor ",
			messages: [][]string{{"este", "es", "un", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 9: Minimo desfasado",
			messages: [][]string{{"", "", "este", "es", "un", "mensaje"}, {"", "este", "", "un", "mensaje"}, {"", "", "es", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			result, err := GetMessage(test.messages...)

			if result != test.result {
				t.Error("Retorna:", result, ",y el valor correcto es:", test.result)
			}

			if !reflect.DeepEqual(err, test.err) {
				t.Error("Retorna error:", err, ",y el valor correcto es:", test.err)
			}
		})
	}
}
