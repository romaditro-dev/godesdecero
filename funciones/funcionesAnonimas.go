package funciones

import "fmt"

//funciones asignadas a variables.
func Calculos() {
	var entero3, entero4 int = 5, 5

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(calculo(10, 9))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * entero3 * entero4
	}

	fmt.Println(calculo(1, 1))
}
