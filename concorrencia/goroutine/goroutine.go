package main

import (
	"fmt"
	"time"
)

func fale(nome, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d) \n", nome, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Porque você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// go fale("Maria", "Ei....", 500)
	// go fale("João", "Opaaa...", 500)

	go fale("Maria", "Entendiii!", 10)
	fale("João", "Parabens!!!", 5)

}
