package main

func inc1(n int) {
	n++
}

// revisão um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n)
	println(n)
}
