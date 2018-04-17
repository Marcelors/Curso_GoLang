package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bm7 struct{}

func (b bm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bm7) fazerBaliza() {
	fmt.Printf("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
