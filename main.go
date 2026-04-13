package main

import (
"encoding/json"
"fmt"
"os"

"financial-expert-system/engine"
"financial-expert-system/knowledge"
"financial-expert-system/report"
)

const benchmarkFile = "data/benchmark.json"

func main() {
// Load industry benchmarks
benchmarks, err := engine.LoadBenchmarks(benchmarkFile)
if err != nil {
fmt.Fprintf(os.Stderr, "Warning: could not load benchmarks: %v\n", err)
}

// Batch mode: go run . --batch data/company.json data/healthy_company.json ...
if len(os.Args) > 1 && os.Args[1] == "--batch" {
RunBatch(os.Args[2:], benchmarks)
return
}

// Single mode: go run . data/company.json
filePath := "data/company.json"
if len(os.Args) > 1 {
filePath = os.Args[1]
}

facts, err := loadFacts(filePath)
if err != nil {
fmt.Fprintf(os.Stderr, "Error: %v\n", err)
os.Exit(1)
}

eng := engine.NewEngine(knowledge.Rules())
result := eng.Run(facts)

// Attach benchmark comparison if industry is specified
if facts.Industry != "" && benchmarks != nil {
result.BenchmarkComparison = engine.Compare(facts, benchmarks)
}

report.Print(facts, result)
}

// loadFacts reads a JSON file and parses it into a FactBase struct.
func loadFacts(path string) (engine.FactBase, error) {
var facts engine.FactBase
data, err := os.ReadFile(path)
if err != nil {
return facts, fmt.Errorf("could not read file: %w", err)
}
if err := json.Unmarshal(data, &facts); err != nil {
return facts, fmt.Errorf("could not parse JSON: %w", err)
}
return facts, nil
}
