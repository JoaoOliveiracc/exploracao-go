package main

import (
	"fmt"
	"time"
)

func maisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle especifica para concorr�ncia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Ainda sem resposta!"
	}
}

func main() {
	campeao := maisRapido(
		"https://www.amazon.com",
		"https://youtube.com",
		"https://google.com",
	)
	fmt.Println(campeao)
}
