package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivo).... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)

	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	// boolean

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Olá
	meu
	nome
	é 
	Higor`
	fmt.Println("o tamanho da string é", len(s2))

	// char ??
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
