// Indicamos que hace parte del paquete workermetric, y será donde definimos la estructura para los mensajes que produciran los Workers como metricas (Cuantos datos procesó cada Worker)
package workermetric

type WorkerMetric struct {
	WorkerID int
	Count    int
}
