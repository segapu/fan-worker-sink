// Indicamos que hace parte del paquete UseCase, y será la clase que interactuará y dependerá de las interfaces para los logs de metricas de los Workers
package usecase

//Importaciones:
//	workermetric: Se importa para poder recibir los mensajes de la cola de metricas
// 	workermetricRepository: Se importa para poder interactuar con la interface
import (
	workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
	workermetricRepository "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric/repository"
)

// Hace el llamado al metodo de la interface y recibe el canal donde están los mensajes que escribirá en consola como Log y define qué hara, más la implementación del cómo hacerlo está en los adapters
func WriteMetricLog(logger workermetricRepository.IWorkerMetrics, message <-chan workermetric.WorkerMetric) {
	logger.EscribirLog(message)
}
