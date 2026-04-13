package main

import (
"encoding/json"
"fmt"
"os"

"financial-expert-system/engine"
"financial-expert-system/knowledge"
)

const batchLine = "════════════════════════════════════════════════════════════════════════════════"
const batchThin = "────────────────────────────────────────────────────────────────────────────────"

// RunBatch analyses multiple company JSON files and prints a comparison table.
func RunBatch(files []string, benchmarks map[string]engine.Benchmark) {
type row struct {
name    string
verdict string
de      float64
margin  float64
roe     float64
growth  float64
cr      float64
ic      float64
industry string
}

var rows []row
eng := engine.NewEngine(knowledge.Rules())

for _, f := range files {
data, err := os.ReadFile(f)
if err != nil {
fmt.Fprintf(os.Stderr, "Skipping %s: %v\n", f, err)
continue
}
var facts engine.FactBase
if err := json.Unmarshal(data, &facts); err != nil {
fmt.Fprintf(os.Stderr, "Skipping %s: invalid JSON\n", f)
continue
}
result := eng.Run(facts)

// Attach benchmark comparison if industry is specified
if facts.Industry != "" && benchmarks != nil {
result.BenchmarkComparison = engine.Compare(facts, benchmarks)
}

rows = append(rows, row{
name:     facts.CompanyName,
verdict:  result.Verdict,
de:       facts.DebtToEquity,
margin:   facts.NetProfitMargin,
roe:      facts.ROE,
growth:   facts.RevenueGrowth,
cr:       facts.CurrentRatio,
ic:       facts.InterestCoverage,
industry: facts.Industry,
})
}

// Print comparison table
fmt.Println(batchLine)
fmt.Println("  BATCH ANALYSIS — COMPANY COMPARISON")
fmt.Println(batchLine)
fmt.Printf("  %-20s %-12s %-14s %6s %8s %7s %8s %6s %6s\n",
"Company", "Status", "Industry", "D/E", "Margin%", "ROE%", "Growth%", "CR", "IC")
fmt.Println(batchThin)
for _, r := range rows {
industry := r.industry
if industry == "" {
industry = "—"
}
fmt.Printf("  %-20s %-12s %-14s %6.2f %7.1f%% %6.1f%% %7.1f%% %6.2f %6.2fx\n",
r.name, r.verdict, industry, r.de, r.margin, r.roe, r.growth, r.cr, r.ic)
}
fmt.Println(batchLine)
fmt.Println()
}
