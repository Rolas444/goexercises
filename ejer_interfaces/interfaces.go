package ejerinterfaces

import (
	"fmt"

	"github.com/rrq/golang/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un %s, y estoy respirando \n", hu.Sexo())
}
