package flip

import (
	"math"
	"testing"
)

func TestFlip(t *testing.T) {
	numFlips := 100_000
	headsTailsProbability := 0.5
	standardDeviation := math.Sqrt(float64(numFlips) * headsTailsProbability * (1 - headsTailsProbability))

	results := make(map[string]int)
	for i := 0; i < numFlips; i++ {
		results[Flip()] += 1
	}

	t.Run("flips produce heads and tails results", func(t *testing.T) {
		if results["Heads"] == 0 {
			t.Errorf("flip did not produce a 'Heads' result")
		}
		if results["Tails"] == 0 {
			t.Errorf("flip did not produce a 'Tails' result")
		}
	})

	t.Run("flips fall within standard deviation", func(t *testing.T) {
		expectedResults := float64(numFlips / 2)
		minimumFlipResults := int(expectedResults - standardDeviation)
		if results["Heads"] < minimumFlipResults {
			t.Errorf("Expected at least %d out of %d flips to be 'Heads', got %d", minimumFlipResults, numFlips, results["Heads"])
		}
		if results["Tails"] < minimumFlipResults {
			t.Errorf("Expected at least %d out of %d flips to be 'Tails' got %d", minimumFlipResults, numFlips, results["Tails"])
		}
	})
}
