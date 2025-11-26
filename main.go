package main

/*"fmt"
"runtime"

variables "github.com/aaroon94/godesde0/Variables"
*/
import (
	ejercicios "github.com/aaroon94/godesde0/ejercicios"
)

func main() {
	/*
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
		fmt.Printf(": Es un, %T\n", numero)

		//condicionales, primero la asignacion y despeus la condicion
		//os := runtime.GOOS
		fmt.Println("--------Concicional if---------")
		if os := runtime.GOOS; os == "linux" || os == "OS X." {
			fmt.Println("Esto no es windows, es", os)
		} else {
			fmt.Println("Esto es windows")
		}
		fmt.Println("--------Condicional Swich---------")
		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/
	ejercicios.Convertir("101")

}
