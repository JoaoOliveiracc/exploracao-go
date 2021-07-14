package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José":     11325.45,
		"Gabriela": 15456.78,
		"Pedro":    1200.0,
	}

	funcsESalarios["Rafael"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
