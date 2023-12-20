package mapas

import "fmt"

func MostrarMapa() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F.M"
	paises["Argentina"] = "Buenos Aires"
	paises["Peru"] = "Lima"

	//fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barca":       39,
		"Real Madrid": 40,
		"Chelsea":     41,
		"Alemania":    50,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo: %s tiene un puntajde de %d\n", equipo, puntaje)
	}

	delete(campeonato, "Barca")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Alemania"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t\n", puntaje, existe)
}
