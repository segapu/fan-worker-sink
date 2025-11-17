// Indicamos que hace parte del paquete result_adapter, y este archivo será quien implementa el metodo de guardado en CSV para la lista resultante obtenida.

package result_adapter

//Importaciones:
//	encoding/csv: Será el paquete que nos ayudará a escribir en un archivo CSV
//	fmt: Que nos permite imprimir mensajes a través de la consola
// 	os: Es el paquete que nos permitirá interactuar con el sistema operativo y poder generar archivos nuevos
// 	strconv: Para convertir el formato y lograr escribir de manera correcta en el archivo CSV
import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Generamos la estructura que podrá ejecutar estos metodos al inyectarlos durante los llamados de los metodos
type ResultAdapter struct{}

// Recibe 1 parametro
//
//	resultados: La lista de los números calculados ordenada de manera acscendente
func (r ResultAdapter) GuardarResultados(resultados []int) {
	//Crea el archvio en el directorio actual, o devuelve un error en caso de que algo falle durante la creación. Si hay un error imprime el mensaje de error y sale del metodo
	file, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("Error creando CSV:", err)
		return
	}

	//Genera el buffer que escribirá en el archivo
	writer := csv.NewWriter(file)
	//Escribe el header del archivo CSV
	writer.Write([]string{"Result"})

	//Escribe los números de manera ordenada en el archivo generado
	for _, r := range resultados {
		writer.Write([]string{strconv.Itoa(r)})
	}

	//Aplaza la llamada de cerrar el archivo hasta que finalice el metodo por cualquier razón
	defer file.Close()
	//Escribe en el archiv forzando lo que hay en el buffer al salir del metodo
	defer writer.Flush()

}
