package main

import (
	"fmt"

	"github.com/romaditro/godesdecero/ejercicios"
)

func main() {
	/*
		variables.MuestroEnteros()

		variables.RestoVariables()

		estado, texto := variables.ConviertoaTexto(1588)

		fmt.Println(estado)
		fmt.Println(texto)
	*/

	numero, texto := ejercicios.ConvNumerico("100")
	fmt.Println(numero)
	fmt.Println(texto)

	numero, texto = ejercicios.ConvNumerico("*/*//")

	fmt.Println(numero)
	fmt.Println(texto)

}
