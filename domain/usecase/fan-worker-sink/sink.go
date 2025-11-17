// Indicamos que hace parte del paquete fan_worker_sink_usecase, y y será la clase que interactuará y dependerá de las interfaces para hacer de las veces de SINK

package fan_worker_sink_usecase

//Importaciones:
// 	fanworkersink_repository: Se importa para poder interactuar con la interface
import (
	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
)

// Hace el llamado al metodo de la interface y el canal del cual recibirá los resultados y los ordenará. Define qué hara, más la implementación del cómo hacerlo está en los adapters
func StartSink(isink fanworkersink_repository.ISink, in <-chan int) []int {
	return isink.Sink(in)
}
