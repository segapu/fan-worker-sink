// Indicamos que hace parte del paquete workerlogRepository, y será donde definimos la interface, el contrato que deberan cumplir
package workerlogRepository

import "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"

type WorkerLogRepositorySave interface {
	GuardarLog(logs <-chan workerlog.WorkerLog)
}
