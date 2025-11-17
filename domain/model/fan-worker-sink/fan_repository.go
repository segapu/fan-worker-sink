// Indicamos que hace parte del paquete workerlogRepository, y será donde definimos la interface, el contrato que deberan cumplir
package fanworkersink_repository

type IFan interface {
	FanOut(numbers []int, out chan<- int)
}
