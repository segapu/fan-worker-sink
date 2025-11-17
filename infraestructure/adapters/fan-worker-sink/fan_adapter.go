package fanworkersink_adapter

type FanAdapter struct{}

// Recibe 2 parametros
//
//	numbers: Una lista de entero que contiene todos los datos que se deben procesar
//	chan out: Recibe un canal en el cual insertará datos unicamente, y es allí donde iran todos los datos que deben ir procesando los Workers.
//Además de que es de tipo FanAdapter
func (FanAdapter) FanOut(nums []int, out chan<- int) {
	for _, n := range nums {
		out <- n
	}
	close(out)
}
