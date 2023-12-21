package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio Sumar")
	result := operacionesMiddleware(sumar)(5, 5)
	fmt.Println(result)

	fmt.Println("Inicio Restar")
	resulte := operacionesMiddleware(restar)(5, 5)
	fmt.Println(resulte)

	fmt.Println("Inicio Multiplicar")
	resulto := operacionesMiddleware(multiplicar)(5, 5)
	fmt.Println(resulto)

}

func operacionesMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operaciones")
		return f(a, b)
	}
}
