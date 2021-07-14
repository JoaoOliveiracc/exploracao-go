package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  1564.45,
			"Gustavo Pereira": 8945.25,
		},
		"J": {
			"João Monteiro": 3000.0,
		},
		"P": {
			"Pedro Filho": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
