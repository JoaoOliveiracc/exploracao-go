package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(3))
	fmt.Println(tipo("teste"))
	fmt.Println(tipo(func() {}))
	fmt.Print(tipo(time.Now()))
}
