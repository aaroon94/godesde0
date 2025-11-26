package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Time time.Time

func RestoVariables() {
	Nombre = "Aaron"
	Estado = true
	Sueldo = 3.1416
	Time = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Time)
}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto

}
