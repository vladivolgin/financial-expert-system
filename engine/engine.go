package engine

// Engine is the forward chaining inference engine.
type Engine struct {
Rules []Rule
}

// NewEngine creates a new inference engine with the given rules.
func NewEngine(rules []Rule) *Engine {
return &Engine{Rules: rules}
}

// Run executes forward chaining until a fixed point is reached.
func (e *Engine) Run(facts FactBase) InferenceResult {
derived := DerivedFacts{}
applied := map[string]bool{}
result := InferenceResult{}

for {
changed := false
for _, rule := range e.Rules {
// Skip rules that have already fired
if applied[rule.Name] {
continue
}
// Evaluate the rule condition
if rule.Condition(facts, derived) {
// Fire the rule — apply its action
rule.Action(&derived)
applied[rule.Name] = true
result.AppliedRules = append(result.AppliedRules, rule.Name)
changed = true
}
}
// Fixed point reached — no new rules fired this pass
if !changed {
break
}
}

// Build the final verdict and recommendations from derived facts
result.DerivedFacts = derived
result.Verdict = resolveVerdict(derived)
result.Recommendations = resolveRecommendations(derived)
return result
}

// resolveVerdict maps DerivedFacts to a human-readable verdict string.
func resolveVerdict(d DerivedFacts) string {
switch {
case d.StatusCritical:
return "🔴 CRITICAL"
case d.StatusWeak:
return "🟠 WEAK"
case d.StatusStable:
return "🟡 STABLE"
case d.StatusStrong:
return "🟢 STRONG"
default:
return "⚪ UNKNOWN"
}
}

// resolveRecommendations builds a list of actionable recommendations based on derived facts.
func resolveRecommendations(d DerivedFacts) []string {
var recs []string
if d.LeverageHigh {
recs = append(recs, "Reduce debt load — D/E ratio is too high")
}
if d.LiquidityLow {
recs = append(recs, "Improve short-term liquidity — risk of default on obligations")
}
if d.ProfitabilityNegative {
recs = append(recs, "Investigate sources of losses — company is unprofitable")
} else if d.ProfitabilityLow {
recs = append(recs, "Review cost structure — profit margin is below threshold")
}
if d.ROELow {
recs = append(recs, "Improve capital efficiency — ROE is below 15%")
}
if d.InterestCoverageWeak {
recs = append(recs, "Critical: interest coverage is dangerously low — risk of default")
}
if d.StatusStrong {
recs = append(recs, "Company is in excellent health — consider strategic expansion")
}
if len(recs) == 0 {
recs = append(recs, "Continue current financial strategy")
}
return recs
}
