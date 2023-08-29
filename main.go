package main

import (
	"fmt"

	"github.com/Fredev01/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(4019)
	fmt.Println(estado)
	fmt.Println(texto)
}
