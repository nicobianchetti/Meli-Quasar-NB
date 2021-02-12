package service

import (
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
			err:      errorInfo,
		},

		{
			name:     "Case 3: Superposición de cadenas",
			messages: [][]string{{"este", "", "", "mensaje", ""}, {"", "es", "uno", "", "secreto"}, {"este", "", "un", "", ""}},
			result:   "",
			err:      errorCadena,
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
			name:     "Case 5.2: Desfasaje 2 - Indeterminación por no tener límites (L/R) ",
			messages: [][]string{{"", "si", "", "es", "", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "este", "", "un", "mensaje", "", ""}},
			result:   "",
			err:      errorCadena,
		},

		{
			name:     "Case 6: Desfasaje 3  ",
			messages: [][]string{{"", "", "", "", "es", "un", ""}, {"este", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 6: Desfasaje 4 - Falta información  ",
			messages: [][]string{{"", "", "", "", "es", "un"}, {"este", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "",
			err:      errorCadena,
		},

		{
			name:     "Case 7: Desfasaje 5  ",
			messages: [][]string{{"este", "", "", "", "es", "un"}, {"", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "",
			err:      errorCadena,
		},

		{
			name:     "Case 8: Desfasaje 6 - Doble referencia array menor ",
			messages: [][]string{{"este", "es", "un", ""}, {"", "", "", ""}, {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 9: Mínimo desfasado",
			messages: [][]string{{"", "", "este", "es", "un", "mensaje"}, {"", "este", "", "un", "mensaje"}, {"", "", "es", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},

		{
			name:     "Case 10: Desfasaje Left/Rigth",
			messages: [][]string{{"", "", "", "es", "", "mensaje", "", ""}, {"este", "es", "un", "mensaje"}, {"", "", "es", "", "mensaje"}},
			result:   "este es un mensaje",
			err:      nil,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			result, err := getMessage(test.messages...)

			if result != test.result {
				t.Error("Retorna:", result, ",y el valor correcto es:", test.result)
			}

			if !reflect.DeepEqual(err, test.err) {
				t.Error("Retorna error:", err, ",y el valor correcto es:", test.err)
			}
		})
	}
}
