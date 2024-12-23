// El paquete principal. Es obligatorio en todo programa ejecutable en Go.
package main

//Importar el paquete "fmt", qu econtiene funciones para formatear e imprimir texto.
import "fmt"

// La función principal del programa.Todo programa en Go comienza ejecutando la funcion main().
func main() {
	// Declaración explicita de variables:
	// "nombre" es una variable de tipo string (cadena de texto), inicializada con el valor "Juan".
	var nombre string = "Juan"
	// "edad" es una variable de tipo int (entero), inicializada con el valor 25.
	var edad int = 25
	// "activo" es una variable de tipo bool (booleano), inicializada con el valor true.
	var activo bool = true

	// Declaración implicita de variables:
	// "apellido" es una variable de tipo string. El tipo infiere automaticamente con el valor "Perez".
	apellido := "Perez"
	// "altura" es una variable de tipo float64 (número decimal). El tipo infiere con el valor 1.75.
	altura := 1.75

	// Imprimir las variables en la consola usando fmt.Println
	// fmt.Println toma uno o más argumentos y los muetra en pantalla, añadiendo un salto de linea.
	fmt.Println("Nombre:", nombre)     // Muestra: Nombre: Juan
	fmt.Println("Apellido:", apellido) // Muestra: Apellido: Perez
	fmt.Println("Edad:", edad)         // Muestra: Edad: 25
	fmt.Println("Altura:", altura)     // Muestra: Altura: 1.75
	fmt.Println("Activo:", activo)     // Muestra: Activo: true
}
