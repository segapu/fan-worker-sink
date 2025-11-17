package fanworkersink_repository

type ISink interface {
	Sink(in <-chan int) []int
}
