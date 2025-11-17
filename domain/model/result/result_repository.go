// Indicamos que hace parte del paquete result_repository, y será donde definimos la interface, el contrato que deberan cumplir
package result_repository

type IResult interface {
	GuardarResultados(resultados []int)
}
