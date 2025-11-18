package workermetric_test

import (
	"testing"

	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

func TestWorkerMetric_Initialization(t *testing.T) {
	expectedWorkerID := 3
	expectedCount := 15

	metric := workermetric.WorkerMetric{
		WorkerID: expectedWorkerID,
		Count:    expectedCount,
	}

	if metric.WorkerID != expectedWorkerID {
		t.Errorf("WorkerID esperado %d, pero se obtuvo %d", expectedWorkerID, metric.WorkerID)
	}

	if metric.Count != expectedCount {
		t.Errorf("Count esperado %d, pero se obtuvo %d", expectedCount, metric.Count)
	}
}
