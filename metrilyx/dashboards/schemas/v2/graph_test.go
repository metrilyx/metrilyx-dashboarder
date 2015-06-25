package v2

import (
	"testing"
)

func Test_NewGraph(t *testing.T) {
	graph := NewGraph(false)
	if graph.Thresholds.Danger.Min != 0 {
		t.Errorf("threshold mismatch: %s", graph.Thresholds)
	}
}
