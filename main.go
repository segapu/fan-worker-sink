// Indicamos que hace parte del paquete Main, y será donde inicia la ejecución de la aplicación.
// Este archivo nos sirve para interactuar como I/O, recibiendo los parametros de ejecicón (--numbers y --workers).
// Adicionalmente tambien nos permitirá conocer cuanto tiempo tardó en realizar la ejecución solicitada.
package main

//Importaciones:
//		flag: 	Que nos permite capturar los datos ingresados como parametros,
// 				además de parsearlos una vez sean recibidos (Ya que si no se hace, tomaría siempre los valores por defecto, sin importar)
//		fmt: 	Que nos permite imprimir mensajes a través de la consola
//		time:  	Que nos permite saber el tiempo de ejecución que tomó
//		orchestrator: Importamos el archivo local que nos permite utilizar las funciones y caracteristicas allí definidas (Para inicializar la lógica en base a los parametros ingresados al momento de ejecutar)
import (
	"flag"
	"fmt"
	"time"

	orchestrator "github.com/sebasgal/fan-worker-sink-go/infraestructure/orchestrator"
)

// El punto donde inicia la aplicación al ejecutarse
func main() {
	//Nos permite conocer el tiempo actual exacto (Año, Mes, Día, Hora, Minuto, Segundo, Milisegundos... etc)
	startTime := time.Now()

	//Parametros opcionales para la ejecución. (--numbers y --workers). Si no recibe ninguno toma por defecto 100 y 3 respectivamente.
	numbersPtr := flag.Int("numbers", 100, "Cantidad de números a procesar")
	workersPtr := flag.Int("workers", 3, "Cantidad de workers a usar")

	//Parseamos los datos recibidos para tomar los valores ingresados
	flag.Parse()

	// Validamos que el parametro --numbers sea mayor que cero, en caso contrario cancelamos la ejecución con un mensaje de error
	if *numbersPtr <= 0 {
		fmt.Println("El parámetro --numbers debe ser mayor que 0")
		return
	}

	// Validamos que el parametro --workers sea mayor que cero, en caso contrario cancelamos la ejecución con un mensaje de error
	if *workersPtr <= 0 {
		fmt.Println("El parámetro --workers debe ser mayor que 0")
		return
	}

	//Llamos el metodo Start que se encuentra dentro del Orchestrator para iniciar la ejecución de la lógica con los parametros recibidos o por defecto y conocer finalmente la cantidad de datos procesados.
	result := orchestrator.Start(*numbersPtr, *workersPtr)
	fmt.Printf("Se procesaron con éxito %d números\n", len(result))
	fmt.Print(result)

	//Conocemos el tiempo que pasó desde que ejecutamos el Start hasta que finalizó la ejecución completa y lo imprimimos en consola.
	totalTime := time.Since(startTime)
	fmt.Printf("El proceso tardó en total: %s\n", totalTime)
}
