package v3

import (
	"testing"
)

func Test_NewPod(t *testing.T) {
	pod := NewPod(false)

	if pod.Panels[0].Type != "line" {
		t.Errorf("graph type mismatch: %s\n", pod)
	}
	pod = NewPod(true)
	if len(pod.Panels) != 0 {
		t.Errorf("skeleton=false failed: %s", pod.Panels)
	}
}
