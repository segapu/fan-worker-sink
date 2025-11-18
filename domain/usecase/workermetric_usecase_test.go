package usecase_test

import (
	"testing"

	workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
	usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase"
)

type mockWorkerMetricLogger struct {
	called bool
	recv   []workermetric.WorkerMetric
}

func (m *mockWorkerMetricLogger) EscribirLog(msgs <-chan workermetric.WorkerMetric) {
	m.called = true

	for metric := range msgs {
		m.recv = append(m.recv, metric)
	}
}


func TestWriteMetricLog_CallsRepository(t *testing.T) {
	mockLogger := &mockWorkerMetricLogger{}

	metrics := make(chan workermetric.WorkerMetric, 2)
	metrics <- workermetric.WorkerMetric{WorkerID: 1, Count: 1}
	metrics <- workermetric.WorkerMetric{WorkerID: 2, Count: 1}
	close(metrics)

	usecase.WriteMetricLog(mockLogger, metrics)

	if !mockLogger.called {
		t.Errorf("Se esperaba que EscribirLog fuera llamado, pero no ocurrió")
	}

	if len(mockLogger.recv) != 2 {
		t.Errorf("Se esperaban 2 métricas, pero se recibieron %d", len(mockLogger.recv))
	}

	if mockLogger.recv[0].WorkerID != 1 || mockLogger.recv[0].Count != 1 {
		t.Errorf("La primera métrica recibida no coincide con la esperada")
	}
	if mockLogger.recv[1].WorkerID != 2 || mockLogger.recv[1].Count != 1 {
		t.Errorf("La segunda métrica recibida no coincide con la esperada")
	}
}
