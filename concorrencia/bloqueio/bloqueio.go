package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação "bloqueante"
	fmt.Println("Depois que for lido")
}

func main() {
	c := make(chan int) // cananl sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação "bloqueante"
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
}
