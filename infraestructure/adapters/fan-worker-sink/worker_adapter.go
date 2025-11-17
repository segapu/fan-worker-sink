package fanworkersink_adapter

import (
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

type WorkerAdapter struct{}

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
