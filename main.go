package main

import (
	"fmt"

	"github.com/rrq/golang/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(155)
	fmt.Println(estado)
	fmt.Println(texto)
}
