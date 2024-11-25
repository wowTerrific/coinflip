package flip

import (
	"slices"
	"testing"
)

func TestFlip(t *testing.T) {
	results := []string{}
	for len(results) < 20 {
		results = append(results, Flip())
	}

	if !slices.Contains(results, "Heads") {
		t.Errorf("flip did not produce a 'Heads' result")
	}
	if !slices.Contains(results, "Tails") {
		t.Errorf("flip did not produce a 'Tails' result")
	}
}
