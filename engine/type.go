package engine

type FactBase struct {
CompanyName      string  `json:"company_name"`
Industry         string  `json:"industry"`
DebtToEquity     float64 `json:"debt_to_equity"`
CurrentRatio     float64 `json:"current_ratio"`
NetProfitMargin  float64 `json:"net_profit_margin"`
RevenueGrowth    float64 `json:"revenue_growth"`
ROE              float64 `json:"roe"`
InterestCoverage float64 `json:"interest_coverage"`
CashRatio        float64 `json:"cash_ratio"`
}

type DerivedFacts struct {
LeverageHigh           bool
LeverageLow            bool
LiquidityGood          bool
LiquidityLow           bool
ProfitabilityGood      bool
ProfitabilityLow       bool
ProfitabilityNegative  bool
ROEStrong              bool
ROELow                 bool
GrowthStrong           bool
GrowthNegative         bool
InterestCoverageStrong bool
InterestCoverageWeak   bool
StatusStrong           bool
StatusStable           bool
StatusWeak             bool
StatusCritical         bool
}

type InferenceResult struct {
AppliedRules        []string
DerivedFacts        DerivedFacts
Verdict             string
Recommendations     []string
BenchmarkComparison *BenchmarkComparison
}
