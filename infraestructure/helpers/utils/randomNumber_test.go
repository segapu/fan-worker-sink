package utils_test

import (
	"testing"

	"github.com/sebasgal/fan-worker-sink-go/infraestructure/helpers/utils"
)

func TestGenerateRandomList_ReturnsCorrectSize(t *testing.T) {
	size := 10
	list := utils.GenerateRandomList(size)

	if len(list) != size {
		t.Fatalf("Se esperaba tamaño %d, pero se obtuvo %d", size, len(list))
	}
}

func TestGenerateRandomList_ValuesWithinRange(t *testing.T) {
	size := 50
	list := utils.GenerateRandomList(size)

	for i, v := range list {
		if v < 0 || v >= 1000 {
			t.Fatalf("Valor fuera de rango en índice %d: %d (esperado: 0 <= n < 1000)", i, v)
		}
	}
}

func TestGenerateRandomList_ProducesDifferentValues(t *testing.T) {
	list1 := utils.GenerateRandomList(20)
	list2 := utils.GenerateRandomList(20)

	areEqual := true
	for i := range list1 {
		if list1[i] != list2[i] {
			areEqual = false
			break
		}
	}

	if areEqual {
		t.Fatalf("Se esperaban listas diferentes, pero ambas fueron iguales")
	}
}
