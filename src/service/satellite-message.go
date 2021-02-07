package service

import (
	"errors"
)

// GetMessage .
// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (string, error) {
	var max int = 0

	for _, v := range messages {
		if len(v) > max {
			max = len(v)
		}
	}

	helper := make([]string, max)

	for _, v := range messages {
		for i, j := range v {
			if j != "" {
				if helper[i] == "" {
					helper[i] = j
				} else {
					if helper[i] != j {
						err := errors.New("Error mensaje. Superposición de cadenas")
						return "", err
					}
				}

			}
		}
	}

	var msg string

	for _, v := range helper {
		if v == "" {
			err := errors.New("Error mensaje. Falta información")
			return "", err
		}
		msg += v + " "
	}

	//Elimino el último espacio
	substring := msg[:(len(msg) - 1)]

	return substring, nil
}
