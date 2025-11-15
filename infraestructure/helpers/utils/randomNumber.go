// Indicamos que hace parte del paquete utils, y será donde generaremos los números aleatorios
package utils

//Importaciones:
//	math/rand: Nos ayudará a poder generar los números aleatorios solicitados
import (
	"math/rand"
)

// Genera los números aleatorios solicitados y los devuelve en una lista de enteros.
// Recibe 1 parametro:
//
//	size: Es la cantidad de números aleatorios que debe generar
//
// Devuelve una lista de enteros con los números aleatorios generados
func GenerateRandomList(size int) []int {
	//Genera una lista del tamaño total de cantidad de números solicitados
	numbers := make([]int, size)
	//En cada posición genera un número aleatorio y lo inyecta al array. El número aleatorio será máximo hasta 1000
	for i := range numbers {
		numbers[i] = rand.Intn(1000)
	}
	return numbers
}
