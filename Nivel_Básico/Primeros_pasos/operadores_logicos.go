// El paquete principal. Es obligatorio en todo programa ejecutable en Go.
package main

//Importar el paquete "fmt", qu econtiene funciones para formatear e imprimir texto.
import "fmt"

// La función principal del programa.Todo programa en Go comienza ejecutando la funcion main().
func main() {
	// Declaración de variables booleanas:
	x := true
	y := false

	// Operaciones lógicos:
	// AND lógico. Devueleve false porque uno de los valores es false.
	resultadoAnd := x && y
	// OR lógico. Devueleve true porque uno de los valores es true.
	resultadoOr := x || y
	// NOT lógico. Devueleve false porque x es true, y se invierte.
	resultadoNotX := !x
	// NOT lógico. Devueleve true porque y es false, y se invierte.
	resultadoNotY := !y

	// Imprime el resultado
	fmt.Println("x && y:", resultadoAnd) // Imprime el resultado de la operación AND entre x y y.
	fmt.Println("x || y:", resultadoOr)  // Imprime el resultado de la operación OR entre x y y.
	fmt.Println("!x:", resultadoNotX)    // Imprime el resultado de la operación NOT sobre x.
	fmt.Println("!y:", resultadoNotY)    // Imprime el resultado de la operación NOT sobre y.
}

/*
Resutlado:

x && y: false
x || y: true
!x: false
!y: true

*/
