package report

import (
"fmt"
"strings"

"financial-expert-system/engine"
)

const line = "════════════════════════════════════════════════════════"
const thin = "────────────────────────────────────────────────────────"

// Print renders the full inference result to stdout.
func Print(f engine.FactBase, r engine.InferenceResult) {
fmt.Println(line)
fmt.Printf("  FINANCIAL EXPERT SYSTEM — %s\n", f.CompanyName)
fmt.Println(line)

printMetrics(f)
printRuleChain(r.AppliedRules)
printDerivedFacts(r.DerivedFacts)
printVerdict(r.Verdict)
printRecommendations(r.Recommendations)

if r.BenchmarkComparison != nil {
printBenchmark(f, r.BenchmarkComparison)
}

fmt.Println(line)
}

// printMetrics displays raw input values with a visual indicator.
func printMetrics(f engine.FactBase) {
fmt.Println("\n📊 INPUT METRICS")
fmt.Println(thin)
fmt.Printf("  %-26s %6.2f   %s\n", "Debt-to-Equity:", f.DebtToEquity, leverageIcon(f.DebtToEquity))
fmt.Printf("  %-26s %6.2f   %s\n", "Current Ratio:", f.CurrentRatio, thresholdIcon(f.CurrentRatio, 1.5, 1.0))
fmt.Printf("  %-26s %6.2f   %s\n", "Cash Ratio:", f.CashRatio, thresholdIcon(f.CashRatio, 0.5, 0.2))
fmt.Printf("  %-26s %5.1f%%   %s\n", "Net Profit Margin:", f.NetProfitMargin, profitIcon(f.NetProfitMargin))
fmt.Printf("  %-26s %5.1f%%   %s\n", "Revenue Growth:", f.RevenueGrowth, thresholdIcon(f.RevenueGrowth, 15.0, 0.0))
fmt.Printf("  %-26s %5.1f%%   %s\n", "ROE:", f.ROE, thresholdIcon(f.ROE, 15.0, 10.0))
fmt.Printf("  %-26s %6.2fx  %s\n", "Interest Coverage:", f.InterestCoverage, thresholdIcon(f.InterestCoverage, 3.0, 1.5))
}

// printRuleChain shows every rule that fired in order — this is the inference chain.
func printRuleChain(rules []string) {
fmt.Println("\n🔗 INFERENCE CHAIN (forward chaining)")
fmt.Println(thin)
for i, name := range rules {
if strings.HasPrefix(name, "Status") && i > 0 {
fmt.Println("  " + thin[:32])
}
fmt.Printf("  %2d. %-30s → fired\n", i+1, name)
}
}

// printDerivedFacts shows only the facts that were derived as true.
func printDerivedFacts(d engine.DerivedFacts) {
fmt.Println("\n📌 DERIVED FACTS")
fmt.Println(thin)
printFact("LeverageHigh", d.LeverageHigh)
printFact("LeverageLow", d.LeverageLow)
printFact("LiquidityGood", d.LiquidityGood)
printFact("LiquidityLow", d.LiquidityLow)
printFact("ProfitabilityGood", d.ProfitabilityGood)
printFact("ProfitabilityLow", d.ProfitabilityLow)
printFact("ProfitabilityNegative", d.ProfitabilityNegative)
printFact("ROEStrong", d.ROEStrong)
printFact("ROELow", d.ROELow)
printFact("GrowthStrong", d.GrowthStrong)
printFact("GrowthNegative", d.GrowthNegative)
printFact("InterestCoverageStrong", d.InterestCoverageStrong)
printFact("InterestCoverageWeak", d.InterestCoverageWeak)
}

// printVerdict prints the final status.
func printVerdict(verdict string) {
fmt.Println("\n" + thin)
fmt.Printf("  OVERALL STATUS: %s\n", verdict)
fmt.Println(thin)
}

// printRecommendations lists all actionable recommendations.
func printRecommendations(recs []string) {
fmt.Println("\n💡 RECOMMENDATIONS")
for _, r := range recs {
fmt.Printf("  • %s\n", r)
}
}

// printBenchmark prints a comparison table between company metrics and industry average.
func printBenchmark(f engine.FactBase, b *engine.BenchmarkComparison) {
fmt.Printf("\n📈 INDUSTRY BENCHMARK — %s\n", strings.ToUpper(b.Industry))
fmt.Println(thin)
fmt.Printf("  %-26s %9s %9s %9s\n", "Metric", "Company", "Industry", "Delta")
fmt.Println(thin)
printBenchmarkRow("Debt-to-Equity",    f.DebtToEquity,    b.AvgDebtToEquity,    b.DeltaDebtToEquity,    true)
printBenchmarkRow("Current Ratio",     f.CurrentRatio,    b.AvgCurrentRatio,    b.DeltaCurrentRatio,    false)
printBenchmarkRow("Cash Ratio",        f.CashRatio,       b.AvgCashRatio,       b.DeltaCashRatio,       false)
printBenchmarkRow("Net Profit Margin", f.NetProfitMargin, b.AvgNetProfitMargin, b.DeltaNetProfitMargin, false)
printBenchmarkRow("Revenue Growth",    f.RevenueGrowth,   b.AvgRevenueGrowth,   b.DeltaRevenueGrowth,   false)
printBenchmarkRow("ROE",               f.ROE,             b.AvgROE,             b.DeltaROE,             false)
printBenchmarkRow("Interest Coverage", f.InterestCoverage,b.AvgInterestCoverage,b.DeltaInterestCoverage,false)
fmt.Println()
}

// printBenchmarkRow prints one row: company value, industry average, delta, and icon.
func printBenchmarkRow(label string, company, industry, delta float64, lowerIsBetter bool) {
var icon string
if lowerIsBetter {
if delta > 0 { icon = "🔴" } else { icon = "✅" }
} else {
if delta >= 0 { icon = "✅" } else { icon = "🔴" }
}
fmt.Printf("  %-26s %9.2f %9.2f %+9.2f  %s\n", label, company, industry, delta, icon)
}

// printFact prints a derived fact only if it is true.
func printFact(name string, value bool) {
if value {
fmt.Printf("  ✔  %s\n", name)
}
}

// leverageIcon returns a risk indicator for debt-to-equity ratio.
func leverageIcon(v float64) string {
if v > 2.0 { return "🔴 HIGH" }
if v > 1.0 { return "⚠️  MODERATE" }
return "✅ LOW"
}

// profitIcon returns a risk indicator for net profit margin.
func profitIcon(v float64) string {
if v < 0   { return "🔴 LOSS" }
if v < 5.0 { return "⚠️  WEAK" }
return "✅ GOOD"
}

// thresholdIcon returns a generic three-level indicator based on two thresholds.
func thresholdIcon(v, good, warn float64) string {
if v >= good { return "✅ GOOD" }
if v >= warn { return "⚠️  BORDERLINE" }
return "🔴 POOR"
}
