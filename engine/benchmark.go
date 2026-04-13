package engine

import (
"encoding/json"
"fmt"
"os"
"strings"
)

// Benchmark holds industry average metrics for comparison.
type Benchmark struct {
DebtToEquity     float64 `json:"debt_to_equity"`
CurrentRatio     float64 `json:"current_ratio"`
NetProfitMargin  float64 `json:"net_profit_margin"`
RevenueGrowth    float64 `json:"revenue_growth"`
ROE              float64 `json:"roe"`
InterestCoverage float64 `json:"interest_coverage"`
CashRatio        float64 `json:"cash_ratio"`
}

// BenchmarkComparison holds both industry averages and deltas.
type BenchmarkComparison struct {
Industry string

// Industry average values
AvgDebtToEquity     float64
AvgCurrentRatio     float64
AvgNetProfitMargin  float64
AvgRevenueGrowth    float64
AvgROE              float64
AvgInterestCoverage float64
AvgCashRatio        float64

// Delta: company value minus industry average
DeltaDebtToEquity     float64
DeltaCurrentRatio     float64
DeltaNetProfitMargin  float64
DeltaRevenueGrowth    float64
DeltaROE              float64
DeltaInterestCoverage float64
DeltaCashRatio        float64
}

// LoadBenchmarks reads the benchmark JSON file and returns a map of industry → Benchmark.
func LoadBenchmarks(path string) (map[string]Benchmark, error) {
data, err := os.ReadFile(path)
if err != nil {
return nil, fmt.Errorf("could not read benchmark file: %w", err)
}
var benchmarks map[string]Benchmark
if err := json.Unmarshal(data, &benchmarks); err != nil {
return nil, fmt.Errorf("could not parse benchmark file: %w", err)
}
return benchmarks, nil
}

// Compare calculates the delta between a company's metrics and the industry average.
// Returns nil if the industry is not found in the benchmark map.
func Compare(facts FactBase, benchmarks map[string]Benchmark) *BenchmarkComparison {
industry := strings.ToLower(facts.Industry)
b, ok := benchmarks[industry]
if !ok {
return nil
}
return &BenchmarkComparison{
Industry: facts.Industry,

// Industry averages
AvgDebtToEquity:     b.DebtToEquity,
AvgCurrentRatio:     b.CurrentRatio,
AvgNetProfitMargin:  b.NetProfitMargin,
AvgRevenueGrowth:    b.RevenueGrowth,
AvgROE:              b.ROE,
AvgInterestCoverage: b.InterestCoverage,
AvgCashRatio:        b.CashRatio,

// Deltas: company - industry
DeltaDebtToEquity:     facts.DebtToEquity - b.DebtToEquity,
DeltaCurrentRatio:     facts.CurrentRatio - b.CurrentRatio,
DeltaNetProfitMargin:  facts.NetProfitMargin - b.NetProfitMargin,
DeltaRevenueGrowth:    facts.RevenueGrowth - b.RevenueGrowth,
DeltaROE:              facts.ROE - b.ROE,
DeltaInterestCoverage: facts.InterestCoverage - b.InterestCoverage,
DeltaCashRatio:        facts.CashRatio - b.CashRatio,
}
}
