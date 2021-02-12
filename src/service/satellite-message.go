package service

import (
	"errors"
)

type messagesSolve struct {
	message [][]string
}

var (
	errorCadena = errors.New("Error mensaje. Superposición de cadenas")
	errorInfo   = errors.New("Error mensaje. Falta información")
	errorIndef  = errors.New("Indefinción")

	maxPhaseShift = 0
)

func message(message chan resultMessage, messages [][]string) {

	defer close(message)

	messageResult, err := getMessage(messages...)

	message <- resultMessage{message: messageResult, err: err}

}

// GetMessage .
// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func getMessage(messages ...[]string) (string, error) {
	lenReference := len(messages[0]) //tomo como referencia el primer elemento para evaluar si todos los mensajes traen la misma cantidad de strings
	var isPhaseShift bool = false
	min := lenReference

	//Verifico las cantidades de strings que tiene cada mensaje
	for _, v := range messages {

		//Al menos un array ya es distinto, hay desfasaje
		if len(v) != lenReference {
			isPhaseShift = true
		}

		if len(v) < min {
			min = len(v)
		}

	}

	if isPhaseShift {
		var messagesMin [][]string
		var msgNormalized [][]string

		//Verificar cantidad de arrays con cantidad mínima de strings
		for _, v := range messages {
			//Si el array tiene la cantidad mínima de strings, apendeo al array de MessagesMin, sino, apendeo a array para normalizar
			if len(v) == min {
				messagesMin = append(messagesMin, v)
			} else {
				msgNormalized = append(msgNormalized, v)
			}
		}

		var arrayReference []string

		//Si tengo mas de un array min, hago merge de ambos para luego determinar ajuste con los restantes
		if len(messagesMin) > 1 {
			result, err := mergeArrays(messagesMin, min)
			if err != nil {
				return "", err
			}

			arrayReference = result

		} else {
			arrayReference = messagesMin[0]
		}

		var arraysNormalizados [][]string

		//Recorro todos los arrays a normalizar por no coincidir con el min de strings
		for _, v := range msgNormalized {

			resultLeft, err := normalizedLeft(arrayReference, v)
			if err != nil {
				return "", err
			}

			if len(resultLeft) > min {
				resultRigth, err := normalizedRight(min, resultLeft)
				if err != nil {
					return "", err
				}

				arraysNormalizados = append(arraysNormalizados, resultRigth)

			} else {
				arraysNormalizados = append(arraysNormalizados, resultLeft)
			}
		}

		arraysNormalizados = append(arraysNormalizados, arrayReference)

		return getMessageProcesed(arraysNormalizados, min)
	}

	return getMessageProcesed(messages, min)

}

func normalizedLeft(reference []string, normalized []string) ([]string, error) {
	var indexReference int
	var indexNormalized int

	// fmt.Println("reference:", reference)

	isMatching := false
	//Recorro array con desfasaje para localizar punto que coincida con otro de array's tomado como referencia
	for i, v := range normalized {
		if v != "" {

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

		}

		if isMatching {
			break
		}
	}

	var phaseShift int

	if isMatching {

		if indexReference == indexNormalized {
			//SI los índices de coincidencias son iguales, entonces devuelve el mismo array para que sea normalizado por límite derecho
			return normalized, nil
		}

		if indexReference > indexNormalized {
			return nil, errorIndef
		}

		//Determinar cuantos lugares debo correr a la izquierda el mensaje desfasado
		phaseShift = indexNormalized - indexReference

		if phaseShift > maxPhaseShift {
			maxPhaseShift = phaseShift
		}

		return normalized[phaseShift:], nil
	}

	//Si no hubo ningún matching de referencia para ver cuantos lugares desplazar a la izquierda, desplazo hasta completar el mínimo
	phaseShift = len(normalized) - len(reference)

	if phaseShift > maxPhaseShift {
		maxPhaseShift = phaseShift
	}

	return normalized[phaseShift:], nil

}

func normalizedRight(min int, normalized []string) ([]string, error) {
	//Recortar array para devolver los primeros <min> elementos. Si la parte recortada queda con un string no vacío, entonces hay un error
	partRigth := normalized[min:]

	for _, v := range partRigth {
		if v != "" {
			return nil, errorCadena
		}
	}
	return normalized[:min], nil
}

//Dado un conjunto de arrays de entrada, se colocan en un hashMap para ver si coinciden las posiciones y se determina el mensaje
func mergeArrays(arrayMessages [][]string, min int) ([]string, error) {
	//Se vuelcan datos de array a un hashMap , si algún caracter no coincide hay un error de superposición de strings diferentes
	helper := make([]string, min)

	for _, v := range arrayMessages {
		for i, j := range v {
			if j != "" {
				//si cadena no es vacía, me fijo si helper ya tenia un valor, sino inserto valor en la posición
				if helper[i] == "" {
					helper[i] = j
				} else {
					if helper[i] != j {
						return nil, errorCadena
					}
				}

			}
		}
	}

	return helper, nil
}

func getMessageProcesed(messages [][]string, min int) (string, error) {

	//Se hace merge de los arrays de messages
	helper, err := mergeArrays(messages, min)
	if err != nil {
		return "", err
	}

	var msg string

	//Verificar que en el merge no haya quedado ningún espacio vacío
	for _, v := range helper {
		if v == "" {
			if maxPhaseShift == 0 {
				return "", errorInfo
			}
			helper = helper[1:]
			maxPhaseShift--
		}
	}

	for _, v := range helper {
		msg += v + " "
	}

	//Elimino el último espacio
	substring := msg[:(len(msg) - 1)]

	return substring, nil
}
