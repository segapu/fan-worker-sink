package fanworkersink_repository

type IFan interface {
	FanOut(numbers []int, out chan<- int)
}
