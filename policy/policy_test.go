package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 86, Capacity: 74, Latency: 25, Risk: 19, Weight: 5}, wantScore: 110, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 90, Capacity: 91, Latency: 15, Risk: 22, Weight: 11}, wantScore: 171, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 104, Capacity: 94, Latency: 12, Risk: 14, Weight: 8}, wantScore: 234, wantDecision: "accept"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
