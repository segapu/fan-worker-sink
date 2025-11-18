# Fan‑Worker‑Sink Go

Este proyecto es una implementación en Go del patron FAN-WORKER-SINK.

Se desarolló utilizando Canales y Go Routines, aprovechando así la fortaleza y narutaleza de Go para tareas paralelas y concurrentes.

Además para su construcciones se planeteó realizarlo con Clean Architecture, para buscar separar el dominio, de los casos de uso, y sus adaptadores, para que así sea más facil su mantenibilidad, además de poder escalar y extenderse más facilmente


## Descripción

El proyecto por defecto genera una lista de 100 números aleatorios que se despachan para procesarse por medio del fan, tendrá 3 workers para procesar cada número, los cuales calculan su cuadrado, y luego el sink los agrupa además de ordenarlos de forma ascendente. 

Para visualizar lo que hace, se capturan e imprimen 2 tipos de logs, unos de metricas por medio de la consola, que nos indicarán cuantos números proceso cada worker, y los otros son logs de traza, que se imprimen por medio de la generación de un archivo .CSV que nos indicará qué número procesó cada worker y qué resultado arrojó para dicho número.


## Arquitectura del Proyecto

![Clen Architecture](https://pitchart.github.io/ddd-cqrs-mvc/resources/img/clean-architecture.png)

El repositorio está organizado de la siguiente forma:

- `domain/model`: Contiene los modelos para WorkerLog y WorkerMetrics que nos dan la estructura necesaria para procesar los dos tipos de logs, y las interfaces para los mismos, además de tambien contener las interfaces para los fan, worker y sink.  
- `domain/usecase/`: Es donde indicamos qué hacer (pero no el cómo) sobre las interfaces.  
- `infraestructure/adapters`: Es donde indicamos cómo se hará el método definido en cada una de las interfaces, es decir, las implementamos
- `infraestructure/helpers/utils`: Se encuentra el archivo utilitario para generar la cantidad de números aleatorios que se le soliciten.
- `infraestructure/orchestrator`: Aquí es donde se orquesta todo el flujo del patron FAN-WORKER-SINK


## Instalación

1. Contar con GO instalado ( Se recomienda seguir la documentación oficial e instalar la version 1.25.4 ). La encuentras [aquí](https://go.dev/doc/install)



2. Clonar el repositorio:

    ```bash
    git clone https://github.com/segapu/fan-worker-sink.git
    cd fan-worker-sink
    ```

## Uso

Una vez se tenga instalado y se encuentre en la ruta .../fan-worker-sink , para ejecutarlo sólo debemos utilizar el comando:

```bash
go run main.go
```

El desarrollo puede recibe 2 parametros que son opcionales:

* numbers: Permite parametrizar la cantidad de números aleatorios que se van a generar y por ende la cantidad que números que se procesarán.
* workers: Permite parametrizar la cantidad de workers que podrán procesar los números que se generen

Por ejemplo para 10.000 números con 5 workers sería

Estos parametros deben ser mayores que 0 ambos, y pueden pasarse en el comando ambos, uno o ninguno

```bash
go run main.go --numbers=10000 --workers=5
```

La ejecución de este desarrollo, nos generará dos archivos .CSV en la raiz de su ejecución

* worker_logs.csv: El cual contendrá la información que cada worker procesó (Qué worker proceso qué número y qué resultado obtuvo). Se visualizará de la siguiente manera, obteniendo la cantidad de registros que números se hayan solicitado

    |Worker ID|Input Number|Result|
    |-|-|-|
    |1|4|16|
    |1|5|25|
    |2|6|36|
    |3|2|4|
    |1|9|81|

* results.csv: El cual contendrá la lista de los números obtenidos ordenados de maner ascendente. Se visualizará de la siguiente manera obteniendo la cantidad de registros que números se hayan solicitado

    |Result|
    |-|
    |4|
    |16|
    |25|
    |36|
    |81|
