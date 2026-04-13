package engine

// Rule represents a single IF → THEN production rule.
type Rule struct {
Name      string
Condition func(facts FactBase, derived DerivedFacts) bool
Action    func(derived *DerivedFacts)
}
