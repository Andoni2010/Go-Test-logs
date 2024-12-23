## Proyecto: Filtrado de Logs en Go

Este proyecto tiene como objetivo demostrar cómo leer y filtrar logs en archivos de texto utilizando Go. El programa lee un archivo de logs (`logs.txt`), filtra las líneas que contienen la palabra "ERROR", y guarda esas líneas en un archivo de salida (`filtered_logs.txt`). Este tipo de filtrado es útil para extraer solo la información relevante de grandes archivos de log.

---

# Proyecto: Filtrado de Logs en Go

Este proyecto tiene como objetivo demostrar cómo leer y filtrar logs en archivos de texto utilizando Go. El programa lee un archivo de logs (`logs.txt`), filtra las líneas que contienen la palabra "ERROR", y guarda esas líneas en un archivo de salida (`filtered_logs.txt`). Este tipo de filtrado es útil para extraer solo la información relevante de grandes archivos de log.

---

## **Descripción del Proyecto**

El programa en Go permite:

1.  Leer un archivo de texto con registros de logs.
2.  Filtrar las líneas que contienen la palabra "ERROR".
3.  Guardar esas líneas en un archivo de salida.

Este proceso es útil para los administradores de sistemas, desarrolladores y equipos de DevOps que necesitan revisar rápidamente los logs para identificar problemas y errores críticos.

---

## Requisitos

-   **Go**: Este proyecto está desarrollado en Go. Asegúrate de tener la última versión de Go instalada en tu máquina.
-   **Git**: Si deseas controlar el código con Git y GitHub, también necesitas tener Git instalado.

---

## Estructura del Proyecto

El proyecto está compuesto por los siguientes archivos:

´´´ 

/log-filter/
  ├── logs.txt               # Archivo de logs de entrada (simulado)
  ├── filtered_logs.txt      # Archivo de salida con los logs filtrados
  ├── log_filter.go          # Código fuente del programa en Go
  └── README.md              # Este archivo

´´´ 

---

## Descripción del Código

El programa realiza las siguientes operaciones:

1.  **Abrir el archivo de logs**: Usa `os.Open` para abrir el archivo `logs.txt`.
2.  **Crear el archivo de salida**: Usa `os.Create` para crear el archivo `filtered_logs.txt`, donde se guardarán las líneas filtradas.
3.  **Leer y filtrar las líneas**: Utiliza `bufio.NewScanner` para leer el archivo línea por línea y filtra las líneas que contienen la palabra "ERROR".
4.  **Escribir el archivo de salida**: Si la línea contiene la palabra "ERROR", se escribe en `filtered_logs.txt`.

---

### **Código principal**

´´´
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
	defer inputFile.Close()
	// Crea el archivo para guardar los logs filtrados
	outputFile, err := os.Create("filtered_logs.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer outputFile.Close()
	// Lee línea por línea y filtra
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
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

´´´ 

## Archivos de Entrada y Salida


### **logs.txt (Archivo de Entrada)**

El archivo `logs.txt` contiene diferentes tipos de registros, incluyendo mensajes de error, info y debug. Un ejemplo del archivo de entrada es el siguiente:

INFO: Application started
ERROR: Failed to connect to database
INFO: User logged in
ERROR: Timeout occurred while fetching data
DEBUG: Cache cleared successfully
ERROR: HOLA BUENAS
DEBUG: Cache cleared successfully

### **filtered_logs.txt (Archivo de Salida)**

El archivo `filtered_logs.txt` es el archivo de salida generado después de ejecutar el programa. Solo contiene las líneas que contienen la palabra "ERROR":

ERROR: Failed to connect to database
ERROR: Timeout occurred while fetching data
ERROR: HOLA BUENAS


## Instrucciones de Uso

### **1. Clona este repositorio**

Si aún no tienes el repositorio, clónalo en tu máquina:

´´´
git clone https://github.com/usuario/repo.git
´´´
### **2. Crea el archivo `logs.txt`**

Crea un archivo llamado `logs.txt` en el mismo directorio que tu programa con el siguiente contenido (o similar):

INFO: Application started
ERROR: Failed to connect to database
INFO: User logged in
ERROR: Timeout occurred while fetching data
DEBUG: Cache cleared successfully
ERROR: HOLA BUENAS
DEBUG: Cache cleared successfully

### **3. Ejecuta el programa**

Compila y ejecuta el código Go con los siguientes comandos:

go run log_filter.go

Este comando leerá el archivo `logs.txt`, filtrará las líneas que contienen "ERROR" y las escribirá en `filtered_logs.txt`.

### **4. Verifica el archivo de salida**

Después de ejecutar el programa, encontrarás un archivo llamado `filtered_logs.txt` con el siguiente contenido:

ERROR: Failed to connect to database
ERROR: Timeout occurred while fetching data
ERROR: HOLA BUENAS

## Comandos Básicos de Git


Aquí te mostramos algunos comandos básicos para manejar tu código en Git:

### **Inicializar un repositorio Git**

`git init` 

Crea un repositorio Git en la carpeta actual.

### **Añadir un repositorio remoto**


`git remote add origin https://github.com/usuario/repo.git`
Vincula tu repositorio local con uno remoto en GitHub.

### **Subir cambios al repositorio remoto**

`git push -u origin main` 

Sube los cambios de la rama `main` al repositorio remoto.

### **Descargar cambios del repositorio remoto**

`git pull origin main` 

Fusiona los cambios del repositorio remoto con tu repositorio local.

## 
## **Pruebas Unitarias**

A continuación, puedes agregar pruebas unitarias para validar el funcionamiento del programa. Por ejemplo, puedes crear un archivo de prueba con contenido simulado para verificar que las líneas que contienen "ERROR" se filtren correctamente.

Ejemplo de prueba básica:

´´´
package main

import (
	"os"
	"testing"
)

func TestFilterLogs(t *testing.T) {
	// Crear archivo de prueba de entrada
	inputFile, err := os.Create("test_logs.txt")
	if err != nil {
		t.Fatalf("error al crear archivo de prueba: %v", err)
	}
	defer inputFile.Close()

	// Agregar contenido de prueba
	inputFile.WriteString("INFO: Application started\n")
	inputFile.WriteString("ERROR: Failed to connect to database\n")

	// Crear archivo de salida
	outputFile, err := os.Create("test_filtered_logs.txt")
	if err != nil {
		t.Fatalf("error al crear archivo de salida: %v", err)
	}
	defer outputFile.Close()

	// Filtrar logs
	err = filterLogs(inputFile, outputFile, "ERROR")
	if err != nil {
		t.Fatalf("error al filtrar logs: %v", err)
	}

	// Verificar el contenido del archivo de salida
	// Implementar aserciones según sea necesario
}

´´´
# 
## **Próximos Pasos**

-   **Ampliar los filtros**: Permitir filtros más complejos, como por fecha o tipo de error.
-   **Agregar pruebas unitarias**: Implementar pruebas automatizadas para validar el código.
-   **Integración en CI/CD**: Integrar el proyecto en una pipeline para procesar logs automáticamente.

## 
## **Conclusión**

Este proyecto demuestra cómo usar Go para procesar logs y filtrar solo las líneas relevantes, como los errores. Es útil para administradores de sistemas, desarrolladores y equipos DevOps que necesitan automatizar el análisis de logs.
