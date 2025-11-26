// Inicio de variables e importacion del paquete
package variables

import (
	"fmt"
)

func MuestroEnteros() {
	var variable int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Variable: ", variable)
	fmt.Println("intde32: ", intde32)
	fmt.Println("intde64: ", intde64)

	fmt.Printf("los valores son %d, %d, %d", variable, intde32, intde64)

}
