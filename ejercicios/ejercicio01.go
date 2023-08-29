package ejercicios

import (
	"strconv"
)

func MostrarValor(valor string) (int, string) {
	valorNumerico, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if valorNumerico > 100 {
		return valorNumerico, "Es mayor a 100"
	} else {
		return valorNumerico, "Es menor a 100"
	}
}
