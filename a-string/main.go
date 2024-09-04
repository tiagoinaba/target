package main

import "fmt"

func aString(s string) int {
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'A' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Print("Insira uma string:\n")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Printf("[Erro] não foi possível ler a string:\n%s", err.Error())
		return
	}
	if n := aString(s); n > 0 {
		fmt.Printf("O caracter 'a' aparece %d vez(es) na string", n)
		return
	}
	fmt.Printf("A string não contém o caracter 'a'")
}
