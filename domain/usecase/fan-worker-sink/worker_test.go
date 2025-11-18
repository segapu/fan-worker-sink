package fan_worker_sink_usecase_test

import (
	"testing"

	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
	fan_worker_sink_usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase/fan-worker-sink"
)

type mockWorker struct {
	called      bool
	recvID      int
	recvFanChan <-chan int
	recvOutChan chan<- int
	recvMetric  chan<- workermetric.WorkerMetric
	recvLog     chan<- workerlog.WorkerLog
}

func (m *mockWorker) Worker(id int,
	fanChan <-chan int,
	workerChan chan<- int,
	metricChan chan<- workermetric.WorkerMetric,
	logChan chan<- workerlog.WorkerLog) {

	m.called = true
	m.recvID = id
	m.recvFanChan = fanChan
	m.recvOutChan = workerChan
	m.recvMetric = metricChan
	m.recvLog = logChan
}


func TestStartWorker_CallsIWorker(t *testing.T) {
	mock := &mockWorker{}

	fanChan := make(chan int)
	workerChan := make(chan int)
	metricChan := make(chan workermetric.WorkerMetric)
	logChan := make(chan workerlog.WorkerLog)

	defer close(fanChan)
	defer close(workerChan)
	defer close(metricChan)
	defer close(logChan)

	workerID := 10

	fan_worker_sink_usecase.StartWorker(
		mock,
		workerID,
		fanChan,
		workerChan,
		metricChan,
		logChan,
	)

	if !mock.called {
		t.Errorf("Se esperaba que Worker fuera llamado, pero no ocurrió")
	}

	if mock.recvID != workerID {
		t.Errorf("Se esperaba ID=%d pero se recibió %d", workerID, mock.recvID)
	}

	if mock.recvFanChan != fanChan {
		t.Errorf("El canal fanChan no coincide")
	}

	if mock.recvOutChan != workerChan {
		t.Errorf("El canal workerChan no coincide")
	}

	if mock.recvMetric != metricChan {
		t.Errorf("El canal metricChan no coincide")
	}

	if mock.recvLog != logChan {
		t.Errorf("El canal logChan no coincide")
	}
}

type mockRealWorker struct{}

func (mockRealWorker) Worker(id int,
	fanChan <-chan int,
	workerChan chan<- int,
	metricChan chan<- workermetric.WorkerMetric,
	logChan chan<- workerlog.WorkerLog) {

	count := 0

	for n := range fanChan {
		squared := n * n
		workerChan <- squared
		count++

		logChan <- workerlog.WorkerLog{
			WorkerID: id,
			Input:    n,
			Output:   squared,
		}
	}

	metricChan <- workermetric.WorkerMetric{
		WorkerID: id,
		Count:    count,
	}
}

func TestStartWorker_ProcessesDataCorrectly(t *testing.T) {
	worker := mockRealWorker{}

	fanChan := make(chan int, 3)
	workerChan := make(chan int, 3)
	metricChan := make(chan workermetric.WorkerMetric, 1)
	logChan := make(chan workerlog.WorkerLog, 3)

	fanChan <- 2
	fanChan <- 3
	fanChan <- 4
	close(fanChan)

	fan_worker_sink_usecase.StartWorker(
		worker,
		1,
		fanChan,
		workerChan,
		metricChan,
		logChan,
	)

	close(workerChan)
	var results []int
	for r := range workerChan {
		results = append(results, r)
	}

	if len(results) != 3 {
		t.Fatalf("Se esperaban 3 resultados, se recibieron %d", len(results))
	}

	if results[0] != 4 || results[1] != 9 || results[2] != 16 {
		t.Errorf("Resultados incorrectos: %v", results)
	}

	close(logChan)
	logs := []workerlog.WorkerLog{}
	for l := range logChan {
		logs = append(logs, l)
	}

	if len(logs) != 3 {
		t.Fatalf("Se esperaban 3 logs, se recibieron %d", len(logs))
	}

	close(metricChan)
	metric := <-metricChan
	if metric.Count != 3 {
		t.Errorf("Se esperaban 3 métricas procesadas, se recibió %d", metric.Count)
	}
}
