package arreglosslices

import "fmt"

var tabla [10]int = [10]int{55, 26, 366, 95}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[9] = 100

	tabla2 := [3]string{"Robert Raul", "Armejo Maque"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
