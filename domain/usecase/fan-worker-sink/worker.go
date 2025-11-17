// Indicamos que hace parte del paquete fan_worker_sink_usecase, y y será la clase que interactuará y dependerá de las interfaces para hacer de las veces de WORKER
package fan_worker_sink_usecase

//Importaciones:
// 	fanworkersink_repository: Se importa para poder interactuar con la interface
// 	workerlog: Se importa para el tipo de datos que será el canal al cual le enviará mensajes de la traza de logs
//	workermetric: Se importa para el tipo de datos que será el canal al cual le enviará mensajes de las metricas de la cantidad de datos procesados por cada worker
import (
	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
	workerlog "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

// Hace el llamado al metodo de la interface y recibe el ID que identifica cada Worker, el canal del cual leerá los números a procesar,
// tambien recibe el canal al cual inyectará los resultados del cuadrado del número calculado, además de los 2 canales de logs, uno de metricas y el otro para
// los logs de lo que procesa exactamente cada worker. Define qué hara, más la implementación del cómo hacerlo está en los adapters
func StartWorker(iworker fanworkersink_repository.IWorker, workerID int,
	fanChan <-chan int,
	workerChan chan<- int,
	metricChan chan<- workermetric.WorkerMetric,
	logChan chan<- workerlog.WorkerLog) {

	iworker.Worker(workerID, fanChan, workerChan, metricChan, logChan)
}
