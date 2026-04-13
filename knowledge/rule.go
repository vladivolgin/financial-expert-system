package knowledge

import "financial-expert-system/engine"

// Rules returns the full knowledge base — production rules in two levels.
// Level 1: raw FactBase → intermediate DerivedFacts
// Level 2: intermediate DerivedFacts → final status
func Rules() []engine.Rule {
return []engine.Rule{

// Level 1: Leverage
{
Name: "HighLeverage",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.DebtToEquity > 2.0
},
Action: func(d *engine.DerivedFacts) { d.LeverageHigh = true },
},
{
Name: "LowLeverage",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.DebtToEquity <= 1.0
},
Action: func(d *engine.DerivedFacts) { d.LeverageLow = true },
},

// Level 1: Liquidity
{
Name: "LiquidityLow",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.CurrentRatio < 1.0
},
Action: func(d *engine.DerivedFacts) { d.LiquidityLow = true },
},
{
Name: "LiquidityGood",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.CurrentRatio >= 1.5
},
Action: func(d *engine.DerivedFacts) { d.LiquidityGood = true },
},

// Level 1: Profitability
{
Name: "ProfitabilityNegative",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.NetProfitMargin < 0
},
Action: func(d *engine.DerivedFacts) {
d.ProfitabilityNegative = true
d.ProfitabilityLow = true
},
},
{
Name: "ProfitabilityLow",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.NetProfitMargin >= 0 && f.NetProfitMargin < 5.0
},
Action: func(d *engine.DerivedFacts) { d.ProfitabilityLow = true },
},
{
Name: "ProfitabilityGood",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.NetProfitMargin >= 5.0
},
Action: func(d *engine.DerivedFacts) { d.ProfitabilityGood = true },
},

// Level 1: ROE
{
Name: "ROELow",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.ROE < 10.0
},
Action: func(d *engine.DerivedFacts) { d.ROELow = true },
},
{
Name: "ROEStrong",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.ROE >= 15.0
},
Action: func(d *engine.DerivedFacts) { d.ROEStrong = true },
},

// Level 1: Growth
{
Name: "GrowthStrong",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.RevenueGrowth >= 15.0
},
Action: func(d *engine.DerivedFacts) { d.GrowthStrong = true },
},
{
Name: "GrowthNegative",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.RevenueGrowth < 0
},
Action: func(d *engine.DerivedFacts) { d.GrowthNegative = true },
},

// Level 1: Interest Coverage
{
Name: "InterestCoverageSafe",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.InterestCoverage >= 3.0
},
Action: func(d *engine.DerivedFacts) { d.InterestCoverageStrong = true },
},
{
Name: "InterestCoverageWeak",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return f.InterestCoverage < 1.5
},
Action: func(d *engine.DerivedFacts) { d.InterestCoverageWeak = true },
},

// Level 2: Final status — depends on Level 1 derived facts
{
Name: "StatusCritical",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return d.LeverageHigh && (d.LiquidityLow || d.ProfitabilityNegative)
},
Action: func(d *engine.DerivedFacts) { d.StatusCritical = true },
},
{
Name: "StatusWeak",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return !d.StatusCritical && (d.LiquidityLow || d.ProfitabilityLow || d.LeverageHigh)
},
Action: func(d *engine.DerivedFacts) { d.StatusWeak = true },
},
{
Name: "StatusStrong",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return d.LiquidityGood && d.ProfitabilityGood && d.LeverageLow && d.ROEStrong
},
Action: func(d *engine.DerivedFacts) { d.StatusStrong = true },
},
{
Name: "StatusStable",
Condition: func(f engine.FactBase, d engine.DerivedFacts) bool {
return !d.StatusCritical && !d.StatusWeak && !d.StatusStrong
},
Action: func(d *engine.DerivedFacts) { d.StatusStable = true },
},
}
}
