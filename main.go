package main

import (
	"fmt"

	variables "github.com/aaroon94/godesde0/Variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	fmt.Println("--------Conversion a Texto con Itoa---------")
	estado, texto := variables.ConviertoTexto(13)
	fmt.Println(estado)
	fmt.Print(texto)
	fmt.Printf(": Es un, %T\n", texto)
	fmt.Println("--------Conversion a Numero con Atoi---------")
	numero := variables.ConviertoNumero("14")
	fmt.Print(numero)
	fmt.Printf(": Es un, %T", numero)

}
