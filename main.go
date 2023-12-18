package main

import (
	"github.com/robertraul/godesde0/variables"
	"fmt"
)

func main(){
	estado, texto := variables.ConviertoaTexto(9555)
	fmt.Println(estado)
	fmt.Println(texto)
}

