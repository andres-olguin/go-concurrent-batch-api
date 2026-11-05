# go-concurrent-batch-api — Ingeniería, pruebas y escalabilidad realista

**Estado:** plan técnico documentado el 04-11-2026. No constituye resultado experimental ejecutado.

## Evidencia de partida

API de tareas Go/Gin con simulaciones de operación y persistencia; concurrencia acotada a 8, lotes máximo 100, cliente benchmark separado; go test y detector de carreras aprobados.

## Incrementos de ingeniería

Sustituir simulación por adaptador real opcional y probado, manejar cancelaciones, contexto, backpressure, timeouts y cierre limpio. Añadir httptest y observabilidad sin datos sensibles.

## Estrategia de escalabilidad

Medir throughput con operaciones reales, latencias p95/p99, tareas rechazadas, colas, memoria y goroutines. Comparar worker pools por recursos y costo bajo carga, no por sleeps.

## Riesgos y condición de aprobación

No declarar MongoDB persistente ni rendimiento enterprise con logs simulados. Controlar tamaños máximos HTTP y límites por identidad antes de exposición.

## Protocolo de experimentación y calidad

Cada comparativa debe incluir referencia de código/commit, entorno, versiones, semilla cuando corresponda, entrada reproducible, comando de ejecución, datos brutos, repetición y límites. Separar datos de ejemplo de resultados realmente ejecutados. Mantener código histórico y coautoría intactos; publicar evidencias de tests de CI y pruebas manuales en su contexto, sin inventar métricas.

## Referencias consultadas (2026-11-04)

- [AWS Well-Architected — rendimiento y costo](https://docs.aws.amazon.com/wellarchitected/).
- [NIST SSDF — prácticas seguras de software](https://csrc.nist.gov/pubs/sp/800/218/final).
- [Google SRE — medición y objetivos](https://sre.google/workbook/implementing-slos/).

Las referencias son metodológicas; no corroboran los benchmarks particulares de este repositorio.
