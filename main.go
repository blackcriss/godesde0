package main

import (
	"fmt"

	"github.com/blackcriss/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1020)
	fmt.Println(estado)
	fmt.Println(texto)
}
