package funciones

import "fmt"

func Calculos() {

	var numero3 int = 32

	suma := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println(suma(10, 60))

	suma = func(num1 int, num2 int) int {
		return num1 + num2*numero3
	}
	fmt.Println(suma(10, 60))
}
