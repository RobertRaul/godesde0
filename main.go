package main

import (
	"github.com/robertraul/godesde0/middleware"
)

//"runtime"

func main() {
	/*
		estado, texto := variables.ConviertoaTexto(1000)
		fmt.Println(estado)
		fmt.Println(texto)
		if os := runtime.GOOS; os == "Linux."{
			fmt.Println("Esto no es Windows")
		}else{
			fmt.Println("Esto SI ES Windows", os)
		}

		switch os := runtime.GOOS; os {
		case "windowss":
			fmt.Println("Estpos e s Womdpws")
		case "Darwin":
			fmt.Println("Estpos es Darwin")
		default:
			fmt.Println(os)

		}

		numero, texto := ejercicios.Mayor100("fas9")
		fmt.Println(numero)
		fmt.Println(texto)
	*/

	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClousure()
	//funciones.Exponencia(2)
	//arreglosslices.MuestroArreglos()
	//arreglosslices.MuestroSlices()
	//arreglosslices.Capacidad()
	//mapas.MostrarMapa()
	//users.AltaUsuario()

	// Robert := new(mo.Hombre)
	// it.HumanosRespirando(Robert)

	// Vero := new(mo.Mujer)
	// it.HumanosRespirando(Vero)
	// canal1 := make(chan bool)
	// go goroutines.MiNombreLento("Robert Raul Armejo Maque", canal1)
	// <-canal1
	//webserver.MiWebServer()
	middleware.MiMiddleware()
}
