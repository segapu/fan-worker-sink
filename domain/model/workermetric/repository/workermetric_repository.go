// Indicamos que hace parte del paquete workermetricRepository, y será donde definimos la interface, el contrato que deberan cumplir
package workermetricRepository

import workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"

type IWorkerMetrics interface {
	EscribirLog(message <-chan workermetric.WorkerMetric)
}
