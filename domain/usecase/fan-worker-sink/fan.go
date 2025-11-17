// Indicamos que hace parte del paquete fan_worker_sink, y este archivo será de las veces de FAN en el patron. Encargado de inyectar los datos que se deben procesar
package fan_worker_sink_usecase

import (
	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
)

// Recibe 2 parametros
//
//	numbers: Una lista de entero que contiene todos los datos que se deben procesar
//	chan out: Recibe un canal en el cual insertará datos unicamente, y es allí donde iran todos los datos que deben ir procesando los Workers.
func StartFan(ifan fanworkersink_repository.IFan, numbers []int, out chan<- int) {
	ifan.FanOut(numbers, out)
}
