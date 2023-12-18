package main

import (
	//"github.com/robertraul/godesde0/variables"
	"fmt"
	//"runtime"
	"github.com/robertraul/godesde0/ejercicios"
)

func main(){
	//estado, texto := variables.ConviertoaTexto(1000)
	//fmt.Println(estado)
	//fmt.Println(texto)	
	// if os := runtime.GOOS; os == "Linux."{
	// 	fmt.Println("Esto no es Windows")
	// }else{
	// 	fmt.Println("Esto SI ES Windows", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "windowss":
	// 	fmt.Println("Estpos e s Womdpws")
	// case "Darwin":
	// 	fmt.Println("Estpos es Darwin")
	// default:
	// 	fmt.Println(os)
		
	// }

	numero, texto := ejercicios.Mayor100("fas9")
	fmt.Println(numero)
	fmt.Println(texto)

	
}

