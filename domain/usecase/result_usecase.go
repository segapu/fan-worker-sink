// Indicamos que hace parte del paquete UseCase, y será la clase que interactuará y dependerá de las interfaces para registrar el resultado
package usecase

//Importaciones:
// 	result_repository: Se importa para poder interactuar con la interface
import (
	result_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/result"
)

// Hace el llamado al metodo de la interface y recibe la lista de los números ordenados para guardarla, define qué hara, más la implementación del cómo hacerlo está en los adapters
func SaveResults(result result_repository.IResult, numbers []int) {
	result.GuardarResultados(numbers)
}
