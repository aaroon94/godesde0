package main

import (
	"fmt"

	variables "github.com/aaroon94/godesde0/Variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(13)
	fmt.Println(estado)
	fmt.Println(texto)
}
