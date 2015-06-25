package v3

import (
	"testing"
)

func Test_Thresholds(t *testing.T) {
	_, err := NewThreshold(100, 0)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	_, err = NewThreshold(0, 100)
	if err == nil {
		t.Errorf("max/min check failed\n")
	}

	NewPanelThresholds()
}
