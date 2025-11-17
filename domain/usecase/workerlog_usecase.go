// Indicamos que hace parte del paquete UseCase, y será la clase que interactuará y dependerá de las interfaces para los logs de acciones de los Workers
package usecase

//Importaciones:
//	workerlog: Se importa para poder recibir los mensajes de la cola de metricas
// 	workerlogRepository: Se importa para poder interactuar con la interface
import (
	workerlog "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	workerlogRepository "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog/repository"
)

// Hace el llamado al metodo de la interface y recibe el canal donde están los mensajes que escribirá en CSV como Log y define qué hara, más la implementación del cómo hacerlo está en los adapters
func SaveLogsCSV(saveCSV workerlogRepository.IWorkerLog, logs <-chan workerlog.WorkerLog) {
	saveCSV.GuardarLog(logs)
}
