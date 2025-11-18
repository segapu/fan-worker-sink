package fanworkersink_adapter_test

import (
	"testing"

	fanworkersink_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/fan-worker-sink"
)


func TestFanAdapter_FanOut_SendsAllNumbers(t *testing.T) {
	adapter := fanworkersink_adapter.FanAdapter{}
	numbers := []int{1, 2, 3}

	out := make(chan int)

	go adapter.FanOut(numbers, out)

	var received []int
	for v := range out {
		received = append(received, v)
	}

	if len(received) != len(numbers) {
		t.Fatalf("Se esperaban %d valores, pero se recibieron %d", len(numbers), len(received))
	}

	for i := range numbers {
		if received[i] != numbers[i] {
			t.Errorf("Valor incorrecto: esperado %d, recibido %d", numbers[i], received[i])
		}
	}
}

func TestFanAdapter_FanOut_ClosesChannel(t *testing.T) {
	adapter := fanworkersink_adapter.FanAdapter{}

	numbers := []int{1, 2, 3}
	out := make(chan int, len(numbers))

	adapter.FanOut(numbers, out)


	for range numbers {
		val, ok := <-out
		if !ok {
			t.Fatalf("El canal se cerró antes de tiempo. Valor leído: %v", val)
		}
	}

	_, ok := <-out
	if ok {
		t.Errorf("Se esperaba que el canal estuviera cerrado después de consumir todos los elementos, pero sigue abierto")
	}
}

func TestFanAdapter_FanOut_EmptySlice(t *testing.T) {
	adapter := fanworkersink_adapter.FanAdapter{}

	numbers := []int{}
	out := make(chan int)

	adapter.FanOut(numbers, out)

	_, ok := <-out

	if ok {
		t.Errorf("El canal debería estar cerrado cuando la lista es vacía, pero sigue abierto")
	}
}
