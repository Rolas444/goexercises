package main

import (
	"fmt"
	"runtime"
)

func main() {
	// estado, texto := variables.ConviertoaTexto(155)
	// fmt.Println(estado)
	// fmt.Println(texto)
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("esto no es windows")
	} else {
		fmt.Println("esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Println("%s \n", os)
	}
}
