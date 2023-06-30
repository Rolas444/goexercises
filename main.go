package main

import "github.com/rrq/golang/middleware"

func main() {
	// estado, texto := variables.ConviertoaTexto(155)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("esto no es windows")
	// } else {
	// 	fmt.Println("esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "Linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Println("%s \n", os)
	// }

	// numero, texto := ejercicios.ConvertNum("0L00")
	// fmt.Println(numero)
	// fmt.Println(texto)

	// teclado.IngresaNumeros()
	// iteraciones.Iterar()
	// fmt.Println(ejercicios.TablaMultiplicar())
	// files.SumaTabla()
	// files.LeoArchivo()
	// funciones.Calculos()
	// funciones.LLamarClosure()
	// funciones.Exponencia(3)
	// arreglosslices.MuestroArreglos()
	// arreglosslices.Capacidad()
	// mapas.MostrarMapas()
	// users.AltaUsuario()
	// Pedro := new(modelos.Hombre)
	// ejerinterfaces.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// ejerinterfaces.HumanosRespirando(Maria)
	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanico()
	// canal1 := make(chan bool)
	// go gorutines.MiNombreLento("Rolando Ramirez", canal1)

	// defer func() {
	// 	<-canal1
	// }()
	// fmt.Println("Estoy aqui")
	// var x string
	// fmt.Scanln(&x)

	// webserver.MiWebserver()
	middleware.MiMiddelware()

}
