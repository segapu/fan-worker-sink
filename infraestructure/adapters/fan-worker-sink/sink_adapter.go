package fanworkersink_adapter

import "sort"

type SinkAdapter struct{}

// Recibe 1 parametro
//
//	chan in: Será el canal del cual leerá el resultado generado por los Workers,
//
// Devuelve una lista de enteros. Está lista será el resultado de la agrupación de todos los datos ordenados de forma ascendente
//Además de que es de tipo SinkAdapter
func (SinkAdapter) Sink(in <-chan int) []int {
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
