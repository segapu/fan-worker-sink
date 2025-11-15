// Indicamos que hace parte del paquete workermetric_adapter, y este archivo será quien implementa el metodo de imprimir en consolalos logs de las metricas de los Workers
package workermetric_adapter

//Importaciones:
//		fmt: 	Que nos permite imprimir mensajes a través de la consola
// 		workermetric: Importamos el paquete para hacer uso de la estructura definida para mapear los logs de metricas de cada Worker (Cuantos datos proceso cada Worker)
import (
	"fmt"

	workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

// Generamos la estructura que podrá ejecutar estos metodos al inyectarlos durante los llamados de los metodos
type WorkerMetricAdapter struct{}

// Escribe la información presente en el canal de metrics (Cuantos datos proceso cada worker)
// Recibe 1 parametro
//
//	chan logs: Es el canal que contiene los logs de metricas que generan los Workers al procesar la información
func (w WorkerMetricAdapter) EscribirLog(message <-chan workermetric.WorkerMetric) {
	fmt.Println("==== Cantidad de números procesados por Worker ====")
	//Imprime en consola la información presente en el canal de metricas
	for m := range message {
		fmt.Printf("Worker %d procesó %d números\n", m.WorkerID, m.Count)
	}
	fmt.Println("=======================================")

}
