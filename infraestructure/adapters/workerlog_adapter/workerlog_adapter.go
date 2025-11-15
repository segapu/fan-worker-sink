// Indicamos que hace parte del paquete workerlog_adapter, y este archivo será quien implementa el metodo de guardado en CSV para los logs de los Workers.
package workerlog_adapter

//Importaciones:
//	encoding/csv: Será el paquete que nos ayudará a escribir en un archivo CSV
//	fmt: Que nos permite imprimir mensajes a través de la consola
// 	os: Es el paquete que nos permitirá interactuar con el sistema operativo y poder generar archivos nuevos
import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
)

// Generamos la estructura que podrá ejecutar estos metodos al inyectarlos durante los llamados de los metodos
type WorkerLogAdapter struct{}

// Guarda la información presente en el canal de logs (Quien proceso qué, qué dato recibio y qué generó)
// Recibe 1 parametro
//
//	chan logs: Es el canal que contiene los logs que generan los Workers al procesar la información
func (w WorkerLogAdapter) GuardarLog(logs <-chan workerlog.WorkerLog) {
	//Crea el archvio en el directorio actual, o devuelve un error en caso de que algo falle durante la creación. Si hay un error imprime el mensaje de error y sale del metodo
	file, err := os.Create("worker_logs.csv")
	if err != nil {
		fmt.Println("Error creando CSV:", err)
		return
	}

	//Genera el buffer que escribirá en el archivo
	writer := csv.NewWriter(file)
	//Escribe el header del archivo CSV 
	writer.Write([]string{"Worker ID", "Input Number", "Result"})

	//Escribe lo que hay presente en el canal dentro del archivo CSV
	for log := range logs {
		writer.Write([]string{
			fmt.Sprint(log.WorkerID),
			fmt.Sprint(log.Input),
			fmt.Sprint(log.Output),
		})
	}
	//Aplaza la llamada de cerrar el archivo hasta que finalice el metodo por cualquier razón
	defer file.Close()
	//Escribe en el archiv forzando lo que hay en el buffer al salir del metodo
	defer writer.Flush()
	fmt.Println("Resultado guardado en el archivo: worker_logs.csv")
}
