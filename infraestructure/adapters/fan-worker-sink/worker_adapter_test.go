package fanworkersink_adapter_test

import (
	"testing"
	"time"

	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
	fanworkersink_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/fan-worker-sink"
)

func TestWorkerAdapter_Worker(t *testing.T) {
	worker := fanworkersink_adapter.WorkerAdapter{}

	in := make(chan int)
	out := make(chan int)
	metrics := make(chan workermetric.WorkerMetric)
	logs := make(chan workerlog.WorkerLog)

	input := []int{1, 2, 3}

	go worker.Worker(10, in, out, metrics, logs)

	go func() {
		for _, n := range input {
			in <- n
		}
		close(in)
	}()

	receivedLogs := []workerlog.WorkerLog{}
	go func() {
		for l := range logs {
			receivedLogs = append(receivedLogs, l)
		}
	}()

	receivedOut := []int{}
	go func() {
		for o := range out {
			receivedOut = append(receivedOut, o)
		}
	}()

	var receivedMetric workermetric.WorkerMetric
	doneMetric := make(chan struct{})

	go func() {
		receivedMetric = <-metrics
		close(doneMetric)
	}()

	time.Sleep(50 * time.Millisecond)

	close(logs)
	close(out)

	expectedOut := []int{1, 4, 9}
	if len(receivedOut) != len(expectedOut) {
		t.Fatalf("expected %d outputs, got %d", len(expectedOut), len(receivedOut))
	}
	for i := range expectedOut {
		if expectedOut[i] != receivedOut[i] {
			t.Errorf("expected out[%d] = %d, got %d", i, expectedOut[i], receivedOut[i])
		}
	}

	if len(receivedLogs) != len(input) {
		t.Fatalf("expected %d logs, got %d", len(input), len(receivedLogs))
	}
	for i, log := range receivedLogs {
		if log.WorkerID != 10 {
			t.Errorf("expected WorkerID 10, got %d", log.WorkerID)
		}
		if log.Input != input[i] {
			t.Errorf("expected log input %d, got %d", input[i], log.Input)
		}
		if log.Output != input[i]*input[i] {
			t.Errorf("expected log output %d, got %d", input[i]*input[i], log.Output)
		}
	}

	<-doneMetric

	if receivedMetric.WorkerID != 10 {
		t.Errorf("expected metric WorkerID 10, got %d", receivedMetric.WorkerID)
	}

	if receivedMetric.Count != len(input) {
		t.Errorf("expected metric Count %d, got %d", len(input), receivedMetric.Count)
	}
}
