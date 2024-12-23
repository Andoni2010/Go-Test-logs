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
	b := 5
	// La variable c se declara.
	c := 10

	// Operadores relacionales:
	// Compara si es a igual a b.
	resultadoIgual := a == b
	// Compara si es a distinto de b.
	resultadoDistinto := a != b
	// Compara si a es mayor que b.
	resultadoMayor := a > b
	// Compara si c es menor que b.
	resultadoMenor := c < a
	// Compara si a es mayor o igual que c.
	resultadoMayorIgual := a >= c
	// Compara si b es menor o igual que c.
	resultadoMenorIgual := b <= c

	// Imprime el resultado de las comparaciones
	fmt.Println("a == b:", resultadoIgual)      // Imprime el resultado de comparación a == b
	fmt.Println("a != b:", resultadoDistinto)   // Imprime el resultado de comparación a != b
	fmt.Println("a > b:", resultadoMayor)       // Imprime el resultado de comparación a > b
	fmt.Println("c < a:", resultadoMenor)       // Imprime el resultado de comparaciónc < a
	fmt.Println("a >= c:", resultadoMayorIgual) // Imprime el resultado de comparación a >= c
	fmt.Println("b <= c:", resultadoMenorIgual) // Imprime el resultado de comparación b <= c
}

/*
Resutlado:

a == b: false
a != b: true
a > b: true
c < a: false
a >= c: true
b <= c: true
*/
