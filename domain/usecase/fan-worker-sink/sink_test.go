package fan_worker_sink_usecase_test

import (
	"testing"

	fan_worker_sink_usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase/fan-worker-sink"
)

type mockSink struct {
	called bool
	input  []int
	output []int
}

func (m *mockSink) Sink(in <-chan int) []int {
	m.called = true

	for v := range in {
		m.input = append(m.input, v)
	}

	return m.output
}

func TestStartSink_CallsSink(t *testing.T) {

	mock := &mockSink{}
	mock.output = []int{10, 20, 30}

	in := make(chan int, 3)
	in <- 1
	in <- 2
	in <- 3
	close(in)

	result := fan_worker_sink_usecase.StartSink(mock, in)

	if !mock.called {
		t.Errorf("Se esperaba que Sink() fuera llamado, pero no ocurrió")
	}

	if len(result) != len(mock.output) {
		t.Errorf("Se esperaban %d resultados, pero se obtuvieron %d",
			len(mock.output), len(result))
	}

	for i := range result {
		if result[i] != mock.output[i] {
			t.Errorf("Resultado #%d incorrecto: esperado %d, obtenido %d",
				i, mock.output[i], result[i])
		}
	}
}

func TestStartSink_ReadsChannelValues(t *testing.T) {

	mock := &mockSink{}

	inputNumbers := []int{5, 10, 15}
	mock.output = []int{}

	in := make(chan int, len(inputNumbers))
	for _, n := range inputNumbers {
		in <- n
	}
	close(in)

	fan_worker_sink_usecase.StartSink(mock, in)

	if len(mock.input) != len(inputNumbers) {
		t.Errorf("El Sink debía leer %d valores, pero leyó %d",
			len(inputNumbers), len(mock.input))
	}

	for i := range inputNumbers {
		if inputNumbers[i] != mock.input[i] {
			t.Errorf("Valor leído incorrecto: esperado %d, obtenido %d",
				inputNumbers[i], mock.input[i])
		}
	}
}

func TestStartSink_EmptyChannel(t *testing.T) {

	mock := &mockSink{}
	mock.output = []int{} 

	in := make(chan int)
	close(in)

	result := fan_worker_sink_usecase.StartSink(mock, in)

	if !mock.called {
		t.Errorf("Sink() debe ser llamado incluso si el canal está vacío")
	}

	if len(result) != 0 {
		t.Errorf("Se esperaba un slice vacío, pero se obtuvieron %d elementos", len(result))
	}

	if len(mock.input) != 0 {
		t.Errorf("No se esperaban valores leídos, pero se leyeron %d", len(mock.input))
	}
}

func TestStartSink_ReturnsMockOutput(t *testing.T) {

	mock := &mockSink{}
	mock.output = []int{99, 100, 101}

	in := make(chan int)
	close(in)

	result := fan_worker_sink_usecase.StartSink(mock, in)

	for i := range mock.output {
		if result[i] != mock.output[i] {
			t.Errorf("El resultado devuelto por StartSink no coincide con el del mock en índice %d", i)
		}
	}
}
