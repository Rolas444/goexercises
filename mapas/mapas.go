package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	// fmt.Println(paises)

	paises["mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	camponato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(camponato)

	for equipo, puntaje := range camponato {
		fmt.Printf("Equipo %s, tiene un puntaje %d \n", equipo, puntaje)
	}

	delete(camponato, "Real Madrid")
	fmt.Println(camponato)

	puntaje, existe := camponato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equip existe = %t", puntaje, existe)

}
