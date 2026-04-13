## Financial Expert System

A rule-based expert system for financial health assessment, built in Go. Uses forward chaining inference to evaluate company financial metrics against a knowledge base of rules and derive an overall health status — with no machine learning or LLM involved.

### What It Does

The system reads a company's financial metrics from a JSON file, fires a chain of inference rules, derives intermediate facts, and produces:

A classified health status: 🔴 CRITICAL / 🟠 WEAK / 🟡 STABLE / 🟢 STRONG
A full inference chain showing every rule that fired
A set of derived facts added to the working memory
Actionable recommendations based on weaknesses found
An industry benchmark comparison with deltas per metric

### How It Works — Forward Chaining

The engine implements a classic forward chaining inference loop:

Working memory is seeded with raw input facts (metric values)
The engine iterates over all rules in the knowledge base
If a rule's conditions are satisfied by current facts → the rule fires, adding new facts to working memory
The loop repeats until no new rules fire (fixed point)
Final status rules check accumulated facts and derive the overall verdict

### This is a two-level inference architecture:

Level 1 rules — evaluate individual metrics (e.g. DebtToEquity > 2.0 → assert LeverageHigh)
Level 2 rules — aggregate Level 1 facts to derive overall status (e.g. LeverageHigh ∧ LiquidityLow ∧ ProfitabilityNegative → assert StatusCritical)

Project Structure
```
financial-expert-system/
├── main.go                  # Entry point, CLI argument parsing
├── go.mod                   # Go module definition
│
├── engine/
│   ├── types.go             # Core types: Company, Fact, WorkingMemory
│   ├── rule.go              # Rule struct and condition/action types
│   └── engine.go            # Forward chaining inference engine
│       └── benchmark.go     # Industry benchmark loading and comparison
│
├── knowledge/
│   └── rules.go             # Full knowledge base — all L1 and L2 rules
│
├── report/
│   └── print.go             # Terminal report formatter
│
└── data/
    ├── benchmark.json        # Industry average benchmarks by sector
    ├── company.json          # Acme Corp — CRITICAL example
    ├── healthy\_company.json  # GrowthTech Inc — STRONG example
    ├── stable\_company.json   # SteadyCo Ltd — STABLE example
    └── weak\_company.json     # StrugglingCo — WEAK example
```

### Running the System

Prerequisites
```
Go 1.21 or later
```
Single company analysis
```
go run . data/company.json
```

### Batch comparison of multiple companies
```
go run . --batch data/company.json data/healthy\company.json data/stable\company.json data/weak\_company.json
```

### Input Format

Each company is defined in a JSON file:
```
{
  "name": "Acme Corp",
  "industry": "retail",
  "debt\to\equity": 2.8,
  "current\_ratio": 0.75,
  "cash\_ratio": 0.15,
  "net\profit\margin": -0.045,
  "revenue\_growth": -0.08,
  "roe": -0.062,
  "interest\_coverage": 1.1
}
```

The industry field is matched against data/benchmark.json to load sector-specific averages.

Inference Rules — Knowledge Base

Level 1 — Metric Classification
```
| Rule | Condition | Derived Fact |
|---|---|---|
| HighLeverage | D/E > 2.0 | LeverageHigh |
| LowLeverage | D/E < 0.8 | LeverageLow |
| LiquidityLow | Current Ratio < 1.0 | LiquidityLow |
| LiquidityGood | Current Ratio ≥ 1.5 | LiquidityGood |
| ProfitabilityNegative | Net Margin < 0 | ProfitabilityNegative |
| ProfitabilityGood | Net Margin ≥ 0.08 | ProfitabilityGood |
| ROELow | ROE < 0.10 | ROELow |
| ROEStrong | ROE ≥ 0.18 | ROEStrong |
| GrowthNegative | Revenue Growth < 0 | GrowthNegative |
| GrowthStrong | Revenue Growth ≥ 0.12 | GrowthStrong |
| InterestCoverageWeak | Interest Coverage < 1.5 | InterestCoverageWeak |
| InterestCoverageSafe | Interest Coverage ≥ 3.0 | InterestCoverageStrong |
```
Level 2 — Status Derivation
```
| Rule | Conditions | Status |
|---|---|---|
| StatusCritical | LeverageHigh ∧ LiquidityLow ∧ ProfitabilityNegative | 🔴 CRITICAL |
| StatusWeak | LeverageHigh ∨ (ROELow ∧ GrowthNegative) | 🟠 WEAK |
| StatusStable | ¬LeverageHigh ∧ ¬LiquidityLow ∧ ¬ProfitabilityNegative | 🟡 STABLE |
| StatusStrong | LeverageLow ∧ LiquidityGood ∧ ROEStrong ∧ GrowthStrong | 🟢 STRONG |
```

### Example Output

════════════════════════════════════════════════════════
  FINANCIAL EXPERT SYSTEM — Acme Corp
════════════════════════════════════════════════════════

📊 INPUT METRICS
  Debt-to-Equity:              2.80   🔴 HIGH
  Current Ratio:               0.75   🔴 POOR
  Net Profit Margin:          -4.5%   🔴 LOSS
  ROE:                        -6.2%   🔴 POOR

🔗 INFERENCE CHAIN (forward chaining)
   1. HighLeverage             → fired
   2. LiquidityLow             → fired
   3. ProfitabilityNegative    → fired
   ...
   7. StatusCritical           → fired

  OVERALL STATUS: 🔴 CRITICAL

📈 INDUSTRY BENCHMARK — RETAIL
  Debt-to-Equity    2.80   vs   1.20   Δ +1.60  🔴
  Net Profit Margin -4.50  vs   4.50   Δ -9.00  🔴


### Threshold Rationale

Thresholds are based on widely accepted financial analysis standards:

D/E > 2.0 — above this level, debt servicing risk increases significantly
Current Ratio < 1.0 — liabilities exceed short-term assets; default risk is real
Net Margin < 0 — company is destroying value, not creating it
ROE < 10% — below the typical cost of equity capital
Interest Coverage < 1.5x — earnings barely cover interest payments; distress territory
Revenue Growth < 0 — shrinking top line compounds all other problems

### Tech Stack

Language: Go 1.21
Dependencies: none (standard library only)
Inference: forward chaining, fixed-point iteration
Knowledge representation: Go structs with condition functions and action closures


