package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 86, Capacity: 74, Latency: 25, Risk: 19, Weight: 5}
	if got := Score(signal); got != 110 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 90, Capacity: 91, Latency: 15, Risk: 22, Weight: 11}
	if got := Score(signal); got != 171 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 104, Capacity: 94, Latency: 12, Risk: 14, Weight: 8}
	if got := Score(signal); got != 234 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
}
