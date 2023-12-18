package main

import (
	"github.com/robertraul/godesde0/variables"
	"fmt"
)

func main(){
	estado, texto := variables.ConviertoaTexto(1000)
	fmt.Println(estado)
	fmt.Println(texto)
}

