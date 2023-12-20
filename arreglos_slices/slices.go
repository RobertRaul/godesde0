package arreglosslices

import "fmt"

var tablaSlices []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{52, 63, 98, 5, 2, 6, 5, 3, 55, 57}

func MuestroSlices() {
	fmt.Println(tablaSlices)
	porcion := arreglo[3:5]
	porcion2 := arreglo[:6]
	porcion3 := arreglo[7:]
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}
func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
