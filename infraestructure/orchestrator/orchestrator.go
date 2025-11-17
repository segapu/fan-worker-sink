// Indicamos que hace parte del paquete Orchestrator. Será el encargado de orquestar el flujo necesario para ejecutar el proyecto con exito en base a los parametros recibidos
package orchestrator

//Importaciones:
//		sync: Lo utilizamos para esperar a que todas las Go Routines indicadas finalicen por completo antes de continuar con la ejecución en el hilo principal
//		workerlog: Importamos el paquete para hacer uso de la estructura definida para mapear los logs de cada Worker (Qué Worker recibio la petición, Qué número recibió y qué resultado arrojó)
// 		workermetric: Importamos el paquete para hacer uso de la estructura definida para mapear los logs de metricas de cada Worker (Cuantos datos proceso cada Worker)
//		usecase: Importamos el paquete para hacer uso de las funcionalidades UseCase definidas para los logs de los Worker (Imprimir en consola y guardar en CSV)
// 		fan_worker_sink: Importamos el paquete donde se encuentra definido los archivos FAN - WORKER - SINK y poder entregarle a cada uno la información que necesita para cumplir el patron
//		workerlog_adapter: Importamos el paquete para poder inyecar la implementación correcta que necesita el UseCase para ejecutar el metodo (Para guardar en CSV dada la implementación del metodo)
//		workermetric_adapter: Importamos el paquete para poder inyectar la implementación correcta que necesita el UseCase para ejecuutar el metodo (Para escribir en consola el resumen de metricas dada la implementación del metodo)
//		utils: Importamos el paquete para poder hacer reuso de una funcionalidad generica y reutilizable (Generar los números aletorios solicitados)
import (
	"sync"

	fanworkersink_repository "github.com/sebasgal/fan-worker-sink-go/domain/model/fan-worker-sink"
	workerlog "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
	usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase"
	fan_worker_sink_usecase "github.com/sebasgal/fan-worker-sink-go/domain/usecase/fan-worker-sink"
	fanworkersink_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/fan-worker-sink"
	workerlog_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/workerlog_adapter"
	workermetric_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/workermetric_adapter"

	utils "github.com/sebasgal/fan-worker-sink-go/infraestructure/helpers/utils"
)

// Generamos la cantidad de números aleatorios solicitados y llamamos el metodo principal de orquestación Execute
// Recibe como parametros:
//
//	'howManyNumber' de tipo entero que representa la cantidad de números que debemos procesar y que se generarán de manera aleatoria
//	'howManyWorkers' de tipo entero que representa la cantidad de workers en los cuales se repartirá la carga de trabajo
//
// Devuelve como resultado una lista de enteros, que contiene el cuadrado de cada número procesado.
func Start(howManynumber int, howManyWorkers int) []int {
	numberList := utils.GenerateRandomList(howManynumber)

	fan := fanworkersink_adapter.FanAdapter{}
	worker := fanworkersink_adapter.WorkerAdapter{}
	sink := fanworkersink_adapter.SinkAdapter{}

	return Execute(numberList, howManyWorkers, fan, sink, worker)

}

// Es el metodo de orquestación principal.
// Recibe como parametros:
//
//	'numbers' de tipo lista de entero que representa la lista con los números aleatorios generados
//	'workerCount' de tipo entero que representa la cantidad de workers en los cuales se repartirá la carga de trabajo
func Execute(numbers []int, workerCount int, fan fanworkersink_repository.IFan, sink fanworkersink_repository.ISink, worker fanworkersink_repository.IWorker) []int {

	//Generamos 4 canales. dos sin buffer (fanChan y workerChan) y dos con buffer (metricsChan y logsChan)
	//	fanChan: Será el canal utilizado por FAN  para inyectar los datos que se deben procesar.
	// 	workerChan: Será el canal utilizado por los Workers para dejar el resultado calculado y será leido por SINK para agrupar dicho resultado
	// 	metricsChan: Será el canal utilizado para que cada Worker registre cuantos datos procesó
	// 	logsChan: Será el canal utilizado por cada Worker para registrar quien procesó qué dato y cual resultado calculó
	fanChan := make(chan int)
	workerChan := make(chan int)
	metricsChan := make(chan workermetric.WorkerMetric, workerCount)
	logsChan := make(chan workerlog.WorkerLog, len(numbers))

	//Se ejecuta una Go Routine para generar una ejecución en paralelo con un nuevo hilo con respecto al hilo principal del programa, y que ejecute la funcion del FAN para inyectar en el canal cada dato que se debe procesar
	go fan_worker_sink_usecase.StartFan(fan, numbers, fanChan)

	//Utilizamos sync.WaitGroup para indicar cuantas Go Routines se deben esperar antes de continuar con el flujo de ejecución en el momento que definamos. Esperaremos la misma cantidad de Go Routines como Workers se hayan solicitado
	var wg sync.WaitGroup
	wg.Add(workerCount)

	//Se generan los Workers solicitados para procesar la información, leyendo del canal FAN y registrando en los demás la información generada. Cuando termina su trabajo, indica que finalizó al sync.WaitGroup
	for i := 1; i <= workerCount; i++ {
		workerID := i
		go func() {
			fan_worker_sink_usecase.StartWorker(worker, workerID, fanChan, workerChan, metricsChan, logsChan)
			wg.Done()
		}()
	}

	//Se genera una nueva Go Routine donde una vez que todos los Workers hayan finalizado de procesar, se procese a cerrar los Channels y evitar deadlock por esperar más datos que nunca van a llegar
	go func() {
		wg.Wait()
		close(workerChan)
		close(metricsChan)
		close(logsChan)
	}()

	//Se ejecuta la funcionalidad del SINK para agrupar y ordenar el resultado de los datos calculados
	results := fan_worker_sink_usecase.StartSink(sink, workerChan)

	//Generamos las instancias que tienen la implementación necesario de los metodos de las interfaces que queremos ejecutar (Guardar en CSV y Escribir logs en Consola)
	workerLogger := workerlog_adapter.WorkerLogAdapter{}
	workerMetric := workermetric_adapter.WorkerMetricAdapter{}

	//Generamos una nueva Go routine con el sync.WaitGroup para generar en otro hilo el archivo en CSV sin intervenir con los logs de la consola en el hilo principal
	wg.Add(1)
	go func() {
		usecase.SaveLogsCSV(workerLogger, logsChan)
		wg.Done()
	}()
	usecase.WriteMetricLog(workerMetric, metricsChan)

	wg.Wait()

	//Retornamos la lista con los resultados agrupados y ordenaos por el SINK
	return results
}
