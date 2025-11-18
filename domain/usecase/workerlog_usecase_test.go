package usecase_test

import (
	"testing"

	workerlog "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase"
)

type mockWorkerLogRepository struct {
	called bool
	recv   []workerlog.WorkerLog
}

func (m *mockWorkerLogRepository) GuardarLog(logs <-chan workerlog.WorkerLog) {
	m.called = true

	for log := range logs {
		m.recv = append(m.recv, log)
	}
}


func TestSaveLogsCSV_CallsRepository(t *testing.T) {
	mockRepo := &mockWorkerLogRepository{}

	logs := make(chan workerlog.WorkerLog, 2)
	logs <- workerlog.WorkerLog{WorkerID: 1, Input: 5, Output: 25}
	logs <- workerlog.WorkerLog{WorkerID: 2, Input: 3, Output: 9}
	close(logs)

	usecase.SaveLogsCSV(mockRepo, logs)

	if !mockRepo.called {
		t.Errorf("Se esperaba que GuardarLog fuera llamado, pero no ocurrió")
	}

	if len(mockRepo.recv) != 2 {
		t.Errorf("Se esperaban 2 logs, pero se recibieron %d", len(mockRepo.recv))
	}

	if mockRepo.recv[0].WorkerID != 1 || mockRepo.recv[0].Output != 25 {
		t.Errorf("El primer log recibido no coincide con el esperado")
	}
	if mockRepo.recv[1].WorkerID != 2 || mockRepo.recv[1].Output != 9 {
		t.Errorf("El segundo log recibido no coincide con el esperado")
	}
}
