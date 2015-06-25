package v2

import (
	"testing"
)

func Test_NewPod(t *testing.T) {
	pod := NewPod(false)

	if pod.Graphs[0].Type != "line" {
		t.Errorf("graph type mismatch: %s\n", pod)
		t.FailNow()
	}
}
