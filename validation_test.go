package main

import (
    "strings"
    "testing"
)

func TestValidateBatchTasks(t *testing.T) {
    if err := validateBatchTasks([]string{"a", "b"}); err != nil { t.Fatal(err) }
    invalid := [][]string{nil, {}, {""}, {"   "}, {strings.Repeat("a", 257)}}
    for _, input := range invalid {
        if err := validateBatchTasks(input); err == nil { t.Fatalf("expected error for %#v", input) }
    }
    valid := make([]string, 100)
    for i := range valid { valid[i] = "ok" }
    if err := validateBatchTasks(valid); err != nil { t.Fatal(err) }
    if err := validateBatchTasks(append(valid, "extra")); err == nil { t.Fatal("expected size limit") }
}
