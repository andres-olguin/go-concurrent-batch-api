# Go Concurrent Batch API — Procesamiento concurrente de lotes

**Laboratorio técnico de concurrencia en Go** con API HTTP basada en Gin. Simula operaciones por lotes, utiliza goroutines acotadas para procesar tareas y proporciona un ejecutable independiente de prueba de estrés local.

> **Estado:** API experimental. La función `saveToMongoSimulated` **no inserta registros en MongoDB**: simula trabajo mediante espera y mensajes de consola. No existe autenticación, autorización ni persistencia empresarial validada. No desplegar públicamente.

## Capacidades comprobadas en código

| Capacidad | Implementación |
|---|---|
| Enviar lote de tareas HTTP | `POST /api/v1/process-batch` |
| Validar tamaño del lote | 1 a 100 tareas |
| Limitar cadena de cada tarea | Máximo 256 bytes, no vacía |
| Acotar concurrencia | Hasta ocho tareas simultáneas |
| Simular procesamiento/persistencia | Esperas y mensajes, **sin escritura MongoDB** |
| Generar carga local | Ejecutable separado en `cmd/benchmark/` |

## Arranque local

**Requisitos:** versión de Go compatible con [`go.mod`](go.mod) — actualmente requiere Go 1.26.5 — y dependencias descargables.

```bash
go test ./...
go test -race ./...
go run .
```

En otra terminal, únicamente contra tu entorno local autorizado:

```bash
curl -X POST http://localhost:8080/api/v1/process-batch \
  -H "Content-Type: application/json" \
  -d '{"tasks":["demo uno","demo dos"]}'
```

En PowerShell, puedes enviar JSON con `Invoke-RestMethod` si el uso de comillas de `curl` es diferente.

## Estructura

| Archivo | Función |
|---|---|
| [`main.go`](main.go) | Servidor y procesamiento simulado |
| [`validation.go`](validation.go) | Validaciones de lote |
| [`validation_test.go`](validation_test.go) | Casos normales e inválidos |
| [`cmd/benchmark/main.go`](cmd/benchmark/main.go) | Cliente de carga independiente |
| [`.github/workflows/go-tests.yml`](.github/workflows/go-tests.yml) | `go test` y detector de carreras |
| [`docs/`](docs/) | Evaluación y evolución |

## Cliente de benchmark

```bash
go run ./cmd/benchmark
```

**Precaución:** genera muchas peticiones concurrentes al servidor local; no ejecutarlo contra sistemas ajenos. El benchmark no prueba escrituras MongoDB, duración de transacciones ni capacidad empresarial. El cliente se separó del paquete raíz porque el archivo histórico tenía una segunda función `main` que impedía compilar con `go test ./...`.

## Evolución prevista

Añadir cancelación por contexto, límites de cuerpo HTTP, permisos, rate limiting, colas con backpressure, persistencia real opcional, pruebas de rutas y benchmarks con cargas representativas. Medir latencia p95/p99, errores y uso de recursos antes de escalar workers.

[Plan de ingeniería](docs/ENGINEERING_ROADMAP_2026-11-04.md) · [Estudio de escalabilidad](docs/ARCHITECTURE_SCALABILITY_2026-11-04.md).

**Seguridad:** no introducir tareas con datos confidenciales mientras el sistema siga registrando mensajes y simulando persistencia.
