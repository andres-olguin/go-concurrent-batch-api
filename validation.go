package main

import (
    "fmt"
    "strings"
)

// validateBatchTasks evita consumo ilimitado accidental y entradas vacías.
func validateBatchTasks(tasks []string) error {
    if len(tasks) == 0 || len(tasks) > 100 {
        return fmt.Errorf("el lote debe contener entre 1 y 100 tareas")
    }
    for i, task := range tasks {
        if strings.TrimSpace(task) == "" || len(task) > 256 {
            return fmt.Errorf("tarea %d inválida: se requieren 1..256 bytes no vacíos", i+1)
        }
    }
    return nil
}
