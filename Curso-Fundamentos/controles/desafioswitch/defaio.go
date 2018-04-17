package main

func notaParaConceito(n float64) string {
	switch true {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "A"
	default:
		return "E"
	}
}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(7.8))
	println(notaParaConceito(5.5))
	println(notaParaConceito(4.2))
	println(notaParaConceito(2.0))
}
