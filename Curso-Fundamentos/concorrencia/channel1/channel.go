package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviado dados para o canal (escrita)
	<-ch    // recendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
