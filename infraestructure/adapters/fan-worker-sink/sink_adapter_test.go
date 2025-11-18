package fanworkersink_adapter_test

import (
	"testing"

	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
	fanworkersink_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/fan-worker-sink"
)

func TestSinkAdapter_ImplementsInterface(t *testing.T) {
	var _ fanworkersink_repository.ISink = (*fanworkersink_adapter.SinkAdapter)(nil)
}


func TestSinkAdapter_Sink_ReturnsSortedResults(t *testing.T) {
	adapter := fanworkersink_adapter.SinkAdapter{}

	in := make(chan int, 5)

	in <- 42
	in <- 7
	in <- 100
	in <- 1
	in <- 50
	close(in)

	result := adapter.Sink(in)

	expected := []int{1, 7, 42, 50, 100}

	if len(result) != len(expected) {
		t.Fatalf("Se esperaban %d resultados, pero se recibieron %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Valor incorrecto en posición %d. Esperado %d, recibido %d",
				i, expected[i], result[i])
		}
	}
}


func TestSinkAdapter_Sink_EmptyChannelReturnsEmptySlice(t *testing.T) {
	adapter := fanworkersink_adapter.SinkAdapter{}

	in := make(chan int)
	close(in)

	result := adapter.Sink(in)

	if len(result) != 0 {
		t.Fatalf("Se esperaba slice vacío, pero se recibieron %v", result)
	}
}

func TestSinkAdapter_Sink_SingleElement(t *testing.T) {
	adapter := fanworkersink_adapter.SinkAdapter{}

	in := make(chan int, 1)
	in <- 99
	close(in)

	result := adapter.Sink(in)

	if len(result) != 1 {
		t.Fatalf("Se esperaba un solo elemento, pero se recibieron %d", len(result))
	}

	if result[0] != 99 {
		t.Errorf("Valor incorrecto. Esperado 99, recibido %d", result[0])
	}
}
