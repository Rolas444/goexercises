package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var err error
var texto string

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("ingrese numero:")
		if scanner.Scan() {
			number1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", number1, i, i*number1)
	}

	return texto

}
