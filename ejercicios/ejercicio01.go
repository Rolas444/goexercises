package ejercicios

import "strconv"

func ConvertNum(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
