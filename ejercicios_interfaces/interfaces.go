package ejercicios_interfaces

import (
	"fmt"

	"github.com/robertraul/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy una %s, y estoy respirando\n", hu.Sexo())
}
