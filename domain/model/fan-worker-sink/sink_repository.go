// Indicamos que hace parte del paquete workerlogRepository, y será donde definimos la interface, el contrato que deberan cumplir
package fanworkersink_repository

type ISink interface {
	Sink(in <-chan int) []int
}
