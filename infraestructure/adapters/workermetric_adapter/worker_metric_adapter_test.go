package workermetric_adapter_test

import (
	"bytes"
	"os"
	"testing"

	workermetric "github.com/sebasgal/fan-worker-sink-go/domain/model/workermetric"
	workermetric_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/workermetric_adapter"
)

func TestWorkerMetricAdapter_EscribirLog_ImprimeMetricas(t *testing.T) {
	// --- Arrange ---

	// Crear un pipe para capturar la salida estándar
	readPipe, writePipe, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = writePipe
	defer func() { os.Stdout = stdout }()

	var buffer bytes.Buffer

	// Gorutina que copia lo que el adapter imprime hacia nuestro buffer
	done := make(chan struct{})
	go func() {
		_, _ = buffer.ReadFrom(readPipe)
		close(done)
	}()

	adapter := workermetric_adapter.WorkerMetricAdapter{}

	message := make(chan workermetric.WorkerMetric, 2)
	message <- workermetric.WorkerMetric{WorkerID: 1, Count: 5}
	message <- workermetric.WorkerMetric{WorkerID: 2, Count: 3}
	close(message)

	// --- Act ---
	adapter.EscribirLog(message)

	// Cerrar pipe para que la gorutina termine
	writePipe.Close()
	<-done

	// --- Assert ---
	expectedHeader := "==== Cantidad de números procesados por Worker ===="
	expectedLine1 := "Worker 1 procesó 5 números"
	expectedLine2 := "Worker 2 procesó 3 números"
	expectedFooter := "======================================="

	out := buffer.Bytes()

	if !bytes.Contains(out, []byte(expectedHeader)) {
		t.Errorf("No se encontró el header esperado:\n%s", expectedHeader)
	}

	if !bytes.Contains(out, []byte(expectedLine1)) {
		t.Errorf("No se encontró la línea esperada:\n%s", expectedLine1)
	}

	if !bytes.Contains(out, []byte(expectedLine2)) {
		t.Errorf("No se encontró la línea esperada:\n%s", expectedLine2)
	}

	if !bytes.Contains(out, []byte(expectedFooter)) {
		t.Errorf("No se encontró el footer esperado:\n%s", expectedFooter)
	}
}
