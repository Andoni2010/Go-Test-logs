// El paquete principal. Es obligatorio en todo programa ejecutable en Go.
package main

//Importar el paquete "fmt", qu econtiene funciones para formatear e imprimir texto.
import "fmt"

// La función principal del programa.Todo programa en Go comienza ejecutando la funcion main().
func main() {
	// Declaración de variables:
	// La variable a se declara.
	a := 10
	// La variable b se declara.
	b := 3

	// Operaciones aritméticas:
	// Suma de la a y b.
	suma := a + b
	// Resta de la a y b.
	resta := a - b
	// Multiplicación de la a y b.
	multiplicacion := a * b
	// División de la a y b..
	division := a / b
	// Rsiduo de la division de a entre b.
	residuo := a % b

	// Imprime el resultado de las comparaciones
	fmt.Println("Suma:", suma)                     // Imprime el resultado de suma.
	fmt.Println("Resta:", resta)                   // Imprime el resultado de resta.
	fmt.Println("Multiplicación:", multiplicacion) // Imprime el resultado de multiplicacion.
	fmt.Println("División:", division)             // Imprime el resultado de division.
	fmt.Println("Residuo:", residuo)               // Imprime el resultado de residuo.
}

/*
Resutlado:

Suma: 13
Resta: 7
Multiplicación: 30
División: 3
Residuo: 1

*/
