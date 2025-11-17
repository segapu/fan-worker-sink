// Indicamos que hace parte del paquete fan_worker_sink, y este archivo será de las veces de SINK en el patron. Encargado de agrupar el resultado y ordenarlo de forma ascendente

package fan_worker_sink_usecase

//Importaciones
//	'sort': Será necesario para utilizar el metodo de ordenamiento de datos.
import (
	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
)

func StartSink(isink fanworkersink_repository.ISink, in <-chan int) []int {
	return isink.Sink(in)
}
