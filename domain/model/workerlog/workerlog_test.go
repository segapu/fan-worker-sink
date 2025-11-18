package workerlog_test

import (
	"testing"

	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
)

func TestWorkerLog_Initialization(t *testing.T) {
	expectedWorkerID := 5
	expectedInput := 10
	expectedOutput := 20

	log := workerlog.WorkerLog{
		WorkerID: expectedWorkerID,
		Input:    expectedInput,
		Output:   expectedOutput,
	}

	if log.WorkerID != expectedWorkerID {
		t.Errorf("WorkerID esperado %d, obtenido %d", expectedWorkerID, log.WorkerID)
	}

	if log.Input != expectedInput {
		t.Errorf("Input esperado %d, obtenido %d", expectedInput, log.Input)
	}

	if log.Output != expectedOutput {
		t.Errorf("Output esperado %d, obtenido %d", expectedOutput, log.Output)
	}
}
