package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresNumero() {
	scanner := bufio.NewScanner(os.Stdin) //pedido de captuira por pantalla
	//scanner, incializar un objeto
	//NewScanner, escaner de teclado
	//os, porque elegiriemos el estandar input lo cual es el teclado Stdin
	fmt.Println("Ingrese numero 1")
	if scanner.Scan() { //se coloca esta condicional, el cual el motivo es si el usuario no digito nada no entre a la codncion
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//panic es para deterner el proceso, como un break
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}
	fmt.Println("Ingrese numero 2")
	if scanner.Scan() { //se coloca esta condicional, el cual el motivo es si el usuario no digito nada no entre a la codncion
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//panic es para deterner el proceso, como un break
			panic("El dato ingresado es incorrecto " + err.Error())
		}
	}
	fmt.Println("Ingrese leyenda")
	if scanner.Scan() { //se coloca esta condicional, el cual el motivo es si el usuario no digito nada no entre a la codncion
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

}
