package fanworkersink_repository

import (
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	"github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
)

type IWorker interface {
	Worker(workerID int,
		fanChan <-chan int,
		workerChan chan<- int,
		metricChan chan<- workermetric.WorkerMetric,
		logChan chan<- workerlog.WorkerLog)
}
