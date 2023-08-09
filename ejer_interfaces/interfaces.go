package ejer_interfaces

import (
	"fmt"

	"github.com/blackcriss/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
