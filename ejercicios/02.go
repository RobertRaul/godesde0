package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var err error
var numero int
var texto string

func Ejercicio02() string{
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}

	}

	for i := 0; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	return texto
}
