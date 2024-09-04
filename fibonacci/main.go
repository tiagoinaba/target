package main

import (
	"fmt"
)

func isFibonacci(num int) bool {
	if num == 0 {
		return true
	}
	a := 0
	b := 1
	c := 1
	for c < num {
		a = b
		b = c
		c = a + b
	}
	return c == num
}

func main() {
	fmt.Printf("=== Fibonacci ===\n")
	fmt.Printf("Insira um número, e te direi se ele pertence à sequência de Fibonacci:\n")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Printf("[Erro] não foi possível ler o número inserido:\n%s", err.Error())
		return
	}
	fib := isFibonacci(num)
	if fib {
		fmt.Printf("O número %d pertence a sequência de Fibonacci\n", num)
		return
	}
	fmt.Printf("O número %d não pertence a sequência de Fibonacci\n", num)
}
