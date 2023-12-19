package iteraciones

import (
	"fmt"
)

func Iterar() {

	for i := 10; i > 0; i-- {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}
}
