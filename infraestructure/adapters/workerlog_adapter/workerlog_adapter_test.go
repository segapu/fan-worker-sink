package workerlog_adapter_test

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"

	workerlog "github.com/sebasgal/fan-worker-sink-go/domain/model/workerlog"
	workerlog_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/workerlog_adapter"
)

func TestWorkerLogAdapter_GuardarLog_CreaYEscribeCSV(t *testing.T) {

	// Crear directorio temporal
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "worker_logs.csv")

	// Cambiar el working directory dentro del test para que el CSV se genere en tmpDir
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	// Arrange
	adapterInstance := workerlog_adapter.WorkerLogAdapter{}

	logs := make(chan workerlog.WorkerLog, 2)
	logs <- workerlog.WorkerLog{WorkerID: 1, Input: 5, Output: 25}
	logs <- workerlog.WorkerLog{WorkerID: 2, Input: 3, Output: 9}
	close(logs)

	// Act
	adapterInstance.GuardarLog(logs)

	// Assert: validar archivo
	file, err := os.Open(tmpFile)
	if err != nil {
		t.Fatalf("No se pudo abrir el archivo CSV generado: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("No se pudo leer el CSV: %v", err)
	}

	// Validar header
	expectedHeader := []string{"Worker ID", "Input Number", "Result"}

	if len(rows) < 1 {
		t.Fatalf("El archivo CSV está vacío, esperaba header y datos")
	}

	header := rows[0]
	for i := range expectedHeader {
		if header[i] != expectedHeader[i] {
			t.Errorf("Header incorrecto en columna %d: esperado=%s recibido=%s",
				i, expectedHeader[i], header[i])
		}
	}

	// Validar contenido de los logs
	if len(rows) != 3 {
		t.Fatalf("Se esperaban 3 filas (header + 2), se obtuvieron %d", len(rows))
	}

	if rows[1][0] != "1" || rows[1][1] != "5" || rows[1][2] != "25" {
		t.Errorf("Fila 1 incorrecta: %v", rows[1])
	}

	if rows[2][0] != "2" || rows[2][1] != "3" || rows[2][2] != "9" {
		t.Errorf("Fila 2 incorrecta: %v", rows[2])
	}
}
