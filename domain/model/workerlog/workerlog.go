// Indicamos que hace parte del paquete workerlog, y será donde definimos la estructura para los mensajes que produciran los Workers como metricas (Quien procesó qué y cual fue el resultado)
package workerlog

type WorkerLog struct {
	WorkerID int
	Input    int
	Output   int
}
