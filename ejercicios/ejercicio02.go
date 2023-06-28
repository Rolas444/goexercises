package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var err error

func GetNumber() {
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
		fmt.Printf("%d x %d = %d \n", number1, i, i*number1)
	}

}
