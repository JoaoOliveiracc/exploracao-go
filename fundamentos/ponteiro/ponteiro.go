package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // Pega o endereço da variável
	*p++   // Desreferenciando (pegandoo valor)
	i++

	// Go não tem aritimética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
