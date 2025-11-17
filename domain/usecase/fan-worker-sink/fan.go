// Indicamos que hace parte del paquete fan_worker_sink_usecase, y y será la clase que interactuará y dependerá de las interfaces para hacer de las veces de FAN
package fan_worker_sink_usecase

//Importaciones:
// 	fanworkersink_repository: Se importa para poder interactuar con la interface
import (
	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
)

// Hace el llamado al metodo de la interface y recibe la lista números que deberá procesar y el canal donde los inyectará. Define qué hara, más la implementación del cómo hacerlo está en los adapters
func StartFan(ifan fanworkersink_repository.IFan, numbers []int, out chan<- int) {
	ifan.FanOut(numbers, out)
}
