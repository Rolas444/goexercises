package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Ese es el segundo mensaje")
}

func EjemploPanico() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}
