/*
1. Crear un paquete nuevo llamado 'ejercicios'
2. Crear un archivo GO en ese paquete llamado 'ejercicio01.go'
3. Crear una funcion Publica, que devuelva dos valores (int y string) y que reciba de parametro un valor 'string'.
4. El parametro de tipo String recibido debera ser convertido a entero y si el entero es > a 100, el String retornado debe decir  "Es mayor a 100" de lo contrario,
	delvolvera el mensaje "Es menor a 100"
5. El valor numerico entero retornado debe de corresponder al string convertido
6. El Main, llamar a dicho funcion asignandola a dos variables y luego mostrar dichas variables por consola.
7. Grabar, ejecutar y subir todo a GitHub
*/

package ejercicios

import (
	"fmt"
	"strconv"
)

func Convertir(cadena string) (int, string) {
	numero, err := strconv.Atoi(cadena)

	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}

	if numero > 100 {
		fmt.Printf("Es un %T \n", numero)
		return numero, "Es mayor a 100"
	} else {
		fmt.Printf("Es un %T \n", numero)
		return numero, "Es Menor a 100"
	}

}
