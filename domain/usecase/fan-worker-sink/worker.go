// Indicamos que hace parte del paquete fan_worker_sink, y este archivo será de las veces de WORKER en el patron. Encargado de procesar los datos con la lógica solicitada (El cuadrado del número recibido)
package fan_worker_sink_usecase

import (
	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

// Recibe 5 parametros
func StartWorker(iworker fanworkersink_repository.IWorker, workerID int,
	fanChan <-chan int,
	workerChan chan<- int,
	metricChan chan<- workermetric.WorkerMetric,
	logChan chan<- workerlog.WorkerLog) {

	iworker.Worker(workerID, fanChan, workerChan, metricChan, logChan)
}

//
//	id: el ID del Worker que está procesando
//	chan in: El canal del cual unicamente leerá  la información que debe procesar
//	chan out: El canal al cual unicamente inyectará el resultado calculado
//	chan metrics: El canal al cual inyectará cuantos datos proceso el Worker
//	chan logs: El canal al cual irá inyectarlo toda la traza que ejecuta el Worker (Qué dato recibio, Quien lo proceso y Qué resultado obtuvo)
// func Worker(id int, in <-chan int, out chan<- int, metrics chan<- workermetric.WorkerMetric, logs chan<- workerlog.WorkerLog) {

// 	//El contador de datos que ha procesado el Worker
// 	count := 0

// 	//Toma los datos del canal de entrada 'chan in', calcula el cuadrado del número recibido, lo devuelve en el canal de salida 'chan out' y aumenta el contador de datos procesados
// 	for n := range in {
// 		squared := n * n
// 		out <- squared
// 		count++

// 		// Registra a detalle la actividad generada por cada Worker y lo inyecta en el canal de logs (Qué dato recibio, Quien lo proceso y Qué resultado obtuvo)
// 		logs <- workerlog.WorkerLog{
// 			WorkerID: id,
// 			Input:    n,
// 			Output:   squared,
// 		}
// 	}

// 	// Registra a detalle la actividad generada por cada Worker y lo inyecta en el canal de metrics (Cuantos datos proceso el Worker)
// 	metrics <- workermetric.WorkerMetric{
// 		WorkerID: id,
// 		Count:    count,
// 	}
// }
