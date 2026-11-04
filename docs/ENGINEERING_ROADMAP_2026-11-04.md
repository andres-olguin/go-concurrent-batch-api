# Plan técnico verificable — go-concurrent-batch-api

**Fecha:** 2026-11-04. Documento de planificación; no representa pruebas completadas ni aprobación comercial.

## Hallazgos de la revisión

API de lotes en Go/Gin, con goroutines, tiempos simulados y función saveToMongoSimulated que solo imprime y duerme; no guarda en MongoDB real.

## Prioridades

1. Corregir descripción y logs para indicar inequívocamente persistencia simulada; nunca mostrar 'guardado exitosamente' como operación real.
2. Limitar número de tareas por lote, tamaño de payload, concurrencia por worker pool, timeouts y cancelación mediante context.
3. Agregar pruebas con httptest para entradas vacías, inválidas y lotes máximos; benchmark real antes de afirmar alto rendimiento.

## Evidencia para cerrar el plan

Las respuestas identifican claramente datos simulados, cargas grandes se rechazan o limitan, go test -race y pruebas de carga reproducibles quedan registradas.

## Procedimiento

Trabajar en ramas con PR, conservar código original y documentar pruebas realmente ejecutadas (comando, sistema, resultado y fecha). Nunca exponer secretos ni datos de clientes/estudiantes. La apertura de este PR documental no certifica seguridad, disponibilidad o funcionamiento.
