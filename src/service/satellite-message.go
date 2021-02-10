package service

import (
	"errors"
	"fmt"
)

type messagesSolve struct {
	message [][]string
}

func message(message chan resultMessage, messages [][]string) {

	defer close(message)

	messageResult, err := getMessage(messages...)

	message <- resultMessage{message: messageResult, err: err}

}

// GetMessage .
// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func getMessage(messages ...[]string) (string, error) {
	// lenMessage := make([]int, 3)
	lenReference := len(messages[0]) //tomo como referencia el primer elemento para evaluar si todos los mensajes traen la misma cantida de strings
	var isPhaseShift bool = false
	min := lenReference
	var minIndex int

	//Verifico las cantidades de strings que tiene cada mensaje
	for i, v := range messages {

		if len(v) != lenReference {
			isPhaseShift = true
		}

		if len(v) < min {
			min = len(v)
			minIndex = i
		}

	}

	var msgNormalized [][]string
	if isPhaseShift {
		// var arrayNormalize [][]string
		for i, v := range messages {
			if i != minIndex && len(v) != min {
				normalized := normalizeMsgList(messages[minIndex], v)
				// fmt.Println("Normalize:", normalize)
				msgNormalized = append(msgNormalized, normalized)
			}
		}
		msgNormalized = append(msgNormalized, messages[minIndex]) //Agrego array referencia
		return getMessageProcesed(msgNormalized, min)
	}

	return getMessageProcesed(messages, min)

}

func getMessageProcesed(messages [][]string, min int) (string, error) {
	helper := make([]string, min)

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

func normalizeMsgList(reference, normalized []string) []string {

	// var valueReference string
	var indexReference int
	var indexNormalized int
	var referenceIsNull bool //Se coloca true cuando el array reference tiene todos los strings vacíos

	//Recorro array con desfasaje para localizar punto que coincida con otro del array tomado como referencia
	for i, v := range normalized {
		isMatching := false
		if v != "" {
			referenceIsNull = false // me aseguro que minimamente tiene un valor no vacío

			//Recorro el array tomado como referencia
			//si el array no es todo nil, tomo el primer string
			for j, n := range reference {
				if n == v {
					indexReference = j
					indexNormalized = i
					isMatching = true
					break
				}
			}

			if isMatching {
				break
			}
		}

	}

	// fmt.Println("indexReference", indexReference)
	// fmt.Println("indexNormalized", indexNormalized)

	var phaseShift int

	if indexReference == indexNormalized {

	} else {
		if indexReference < indexNormalized {
			phaseShift = indexNormalized - indexReference
		} else {
			fmt.Println("Indefinción")
		}
	}

	//Determinar límites para ver si corto a izq o a derecha
	normalize := normalized[phaseShift:]

	fmt.Println("Normalizado", normalize)
	_ = referenceIsNull

	return normalize
}
