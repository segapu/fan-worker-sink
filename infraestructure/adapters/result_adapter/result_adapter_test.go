package result_adapter_test

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	result_adapter "github.com/sebasgal/fan-worker-sink-go/infraestructure/adapters/result_adapter"
)

func TestResultAdapter_GuardarResultados_CreaYEscribeCSV(t *testing.T) {

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "result.csv")

	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	adapterInstance := result_adapter.ResultAdapter{}
	numbers := []int{1, 4, 9, 16, 25}

	adapterInstance.GuardarResultados(numbers)

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

	expectedHeader := []string{"Result"}

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

	if len(rows) != len(numbers)+1 {
		t.Fatalf("Se esperaban %d filas (header + %d números), se obtuvieron %d", len(numbers)+1, len(numbers), len(rows))
	}

	for i, val := range numbers {
		numStr := strconv.Itoa(val)
		if rows[i+1][0] != numStr {
			t.Errorf("Fila %d incorrecta: esperado=%s recibido=%s", i+1, numStr, rows[i+1][0])
		}
	}
}
