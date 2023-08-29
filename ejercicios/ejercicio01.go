package ejercicios

import (
	"strconv"
)

func MostrarValor(valor string) (int, string) {
	var message string
	valorNumerico, err := strconv.Atoi(valor)
	if err == nil {
		if valorNumerico > 100 {
			message = "Es mayor a 100"
		} else {
			message = "Es menor a 100"
		}
	} else {
		message = "Ocurrio un error"
	}
	return valorNumerico, message
}
