package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abre el archivo de logs
	inputFile, err := os.Open("logs.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer inputFile.Close() // Corrección: `Close` en vez de `Closed`

	// Crea el archivo para guardar los logs filtrados
	outputFile, err := os.Create("filtered_logs.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer outputFile.Close() // Corrección: `Close` en vez de `Closed`

	// Lee línea por línea y filtra
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			_, err := outputFile.WriteString(line + "\n") // Corrección: `WriteString` en vez de `WriteStrings`
			if err != nil {                               // Corrección: eliminar `:=` en el if
				fmt.Println("Error al escribir en el archivo:", err)
				return
			}
		}
	}

	// Verifica si hubo algún error durante el escaneo
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}
