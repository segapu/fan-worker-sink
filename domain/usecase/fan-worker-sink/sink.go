// Indicamos que hace parte del paquete fan_worker_sink, y este archivo será de las veces de SINK en el patron. Encargado de agrupar el resultado y ordenarlo de forma ascendente

package fan_worker_sink

//Importaciones
//	'sort': Será necesario para utilizar el metodo de ordenamiento de datos.
import (
	"sort"
)

// Recibe 1 parametro
//
//	chan in: Será el canal del cual leerá el resultado generado por los Workers,
//
// Devuelve una lista de enteros. Está lista será el resultado de la agrupación de todos los datos ordenados de forma ascendente
func Sink(in <-chan int) []int {
	//genera una lista de números que irá incrementando cada que procesa alguno.
	results := make([]int, 0)

	//Adiciona cada número presenste en el channel dentro del la lista del resultado
	for r := range in {
		results = append(results, r)
	}

	//Ordena los valores de manera ascendente y devuelve el resultado
	sort.Ints(results)
	return results
}
