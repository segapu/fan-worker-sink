package fan_worker_sink_usecase_test

import (
	"testing"

	fan_worker_sink_usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase/fan-worker-sink"
)


type mockFan struct {
	called bool
	recv   []int
}

func (m *mockFan) FanOut(numbers []int, out chan<- int) {
	m.called = true

	for _, n := range numbers {
		m.recv = append(m.recv, n)
		out <- n
	}

	close(out)
}

func TestStartFan_CallsFanOut(t *testing.T) {

	mock := &mockFan{}
	numbers := []int{5, 10, 15}
	out := make(chan int, len(numbers))

	fan_worker_sink_usecase.StartFan(mock, numbers, out)

	if !mock.called {
		t.Errorf("Se esperaba que FanOut fuera llamado, pero nunca ocurrió")
	}

	if len(mock.recv) != len(numbers) {
		t.Errorf("FanOut debía recibir %d números, recibió %d",
			len(numbers), len(mock.recv))
	}

	for i := range numbers {
		if mock.recv[i] != numbers[i] {
			t.Errorf("El número recibido #%d no coincide: esperado %d, obtenido %d",
				i, numbers[i], mock.recv[i])
		}
	}
}

func TestStartFan_WritesToOutputChannel(t *testing.T) {

	mock := &mockFan{}
	numbers := []int{1, 2, 3}
	out := make(chan int, len(numbers))

	fan_worker_sink_usecase.StartFan(mock, numbers, out)

	var result []int
	for v := range out {
		result = append(result, v)
	}

	if len(result) != len(numbers) {
		t.Errorf("Se esperaban %d valores en el canal, pero se recibieron %d",
			len(numbers), len(result))
	}

	for i := range numbers {
		if result[i] != numbers[i] {
			t.Errorf("El valor #%d en el canal no coincide: esperado %d, obtenido %d",
				i, numbers[i], result[i])
		}
	}
}

func TestStartFan_WithEmptyNumbers(t *testing.T) {

	mock := &mockFan{}
	numbers := []int{}
	out := make(chan int, 1)

	fan_worker_sink_usecase.StartFan(mock, numbers, out)

	if !mock.called {
		t.Errorf("Incluso si la lista está vacía, la llamada a FanOut debe ocurrir (delegación)")
	}

	val, ok := <-out
	if ok {
		t.Errorf("El canal debía estar cerrado, pero se recibió un valor: %d", val)
	}
}
