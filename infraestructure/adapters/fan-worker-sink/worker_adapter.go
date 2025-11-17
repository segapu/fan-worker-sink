package fanworkersink_adapter

import (
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

type WorkerAdapter struct{}

// Recibe 5 parametros
//
//	id: el ID del Worker que está procesando
//	chan in: El canal del cual unicamente leerá  la información que debe procesar
//	chan out: El canal al cual unicamente inyectará el resultado calculado
//	chan metrics: El canal al cual inyectará cuantos datos proceso el Worker
//	chan logs: El canal al cual irá inyectarlo toda la traza que ejecuta el Worker (Qué dato recibio, Quien lo proceso y Qué resultado obtuvo)
//
// Además de que es de tipo WorkerAdapter
func (WorkerAdapter) Worker(id int, in <-chan int, out chan<- int, metrics chan<- workermetric.WorkerMetric, logs chan<- workerlog.WorkerLog) {

	//El contador de datos que ha procesado el Worker
	count := 0

	//Toma los datos del canal de entrada 'chan in', calcula el cuadrado del número recibido, lo devuelve en el canal de salida 'chan out' y aumenta el contador de datos procesados
	for n := range in {
		squared := n * n
		out <- squared
		count++

		// Registra a detalle la actividad generada por cada Worker y lo inyecta en el canal de logs (Qué dato recibio, Quien lo proceso y Qué resultado obtuvo)
		logs <- workerlog.WorkerLog{
			WorkerID: id,
			Input:    n,
			Output:   squared,
		}
	}

	// Registra a detalle la actividad generada por cada Worker y lo inyecta en el canal de metrics (Cuantos datos proceso el Worker)
	metrics <- workermetric.WorkerMetric{
		WorkerID: id,
		Count:    count,
	}
}
