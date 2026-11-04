# Go Concurrent Batch API — laboratorio experimental

API de procesamiento concurrente por lotes desarrollada con Go y Gin. El ejemplo usa tareas simuladas y IDs BSON para demostrar concurrencia y métricas de latencia.

**Importante:** `saveToMongoSimulated` NO inserta en MongoDB: espera 30 ms e imprime un mensaje de demostración. No hay persistencia real, autenticación ni autorización verificadas. **No exponer el servidor a Internet.**

## Uso

```bash
go test ./...
go test -race ./...
go run .
```

Ejemplo de petición local:

```bash
curl -X POST http://localhost:8080/api/v1/process-batch -H "Content-Type: application/json" -d '{"tasks":["demo uno","demo dos"]}'
```

El experimento limita lotes a 100 tareas, textos de hasta 256 bytes por tarea y un máximo de ocho workers simultáneos. Los resultados no tienen orden garantizado por el procesamiento paralelo. Estos límites son una medida inicial, no una certificación anti-DoS.

## Próximos hitos

Autenticación y autorización, cancelación de contexto, rate limiting, límites HTTP, cachés, timeouts, persistencia opt-in y pruebas de carga verificables. Consulta `docs/ENGINEERING_ROADMAP_2026-11-04.md`.

## Datos y seguridad

Usar exclusivamente tareas de ejemplo. No registrar información confidencial en logs, no confiar en los identificadores de Mongo como señal de persistencia ni prometer rendimiento por tiempos simulados.


## Cliente de benchmark aislado

El archivo histórico `benchmark.go` tenía otra función `main` y una definición duplicada de `BatchRequest`, por lo que `go test ./...` fallaba al compilar el paquete raíz. Se conserva su contenido en `cmd/benchmark/main.go` como ejecutable separado.

Ejecutar primero el servidor local y luego, solo contra tu entorno autorizado: `go run ./cmd/benchmark`. El cliente genera carga concurrente y **no debe dirigirse a servidores externos sin permiso**. Sus tiempos no constituyen un benchmark de persistencia MongoDB real.
