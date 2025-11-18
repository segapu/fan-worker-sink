package usecase_test

import (
	"testing"

	usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase"
)

type mockResultRepository struct {
	called   bool
	received []int
}

func (m *mockResultRepository) GuardarResultados(numbers []int) {
	m.called = true
	m.received = append(m.received, numbers...)
}


func TestSaveResults_CallsRepository(t *testing.T) {

	mockRepo := &mockResultRepository{}
	numbers := []int{1, 4, 9, 16, 25}

	usecase.SaveResults(mockRepo, numbers)

	if !mockRepo.called {
		t.Errorf("Se esperaba que GuardarResultados fuera llamado, pero no ocurrió")
	}

	if len(mockRepo.received) != len(numbers) {
		t.Errorf("Se esperaban %d resultados, pero se recibieron %d", len(numbers), len(mockRepo.received))
	}

	for i, val := range numbers {
		if mockRepo.received[i] != val {
			t.Errorf("Resultado en posición %d incorrecto: got %d, want %d", i, mockRepo.received[i], val)
		}
	}
}
