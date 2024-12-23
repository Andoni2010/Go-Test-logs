// El paquete principal. Es obligatorio en todo programa ejecutable en Go.
package main

//Importar el paquete "fmt", qu econtiene funciones para formatear e imprimir texto.
import "fmt"

// La función principal del programa.Todo programa en Go comienza ejecutando la funcion main().
func main() {
	// llama a la función Println del paquete fmt para imprimir texto en la consola
	// Aquí imprime "Hola mundo" seguido de un salto de línea.
	fmt.Println("Hola mundo")
}

// Para ejecutar el programa es con : go run [nombre del archivo].go
// Para formatear el código y garantizar que siga las convenciones oficiales de Go : go fmt [nombre del archivo].go
