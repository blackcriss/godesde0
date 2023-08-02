package ejercicios

import (
	"strconv"
)

func Cien(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Hubo un error // " + err.Error()
	}

	if numero > 100 {
		return numero, "El numero es mayor a 100"
	} else {
		return numero, "El numero es igual o menor a 100"
	}
}
