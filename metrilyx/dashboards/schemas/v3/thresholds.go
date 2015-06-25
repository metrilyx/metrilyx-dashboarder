package v3

import (
	"fmt"
)

type Threshold struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

func NewThreshold(max, min float64) (Threshold, error) {
	var t Threshold
	if max < min {
		return t, fmt.Errorf("max cannot be less than min")
	}
	t = Threshold{max, min}
	return t, nil
}

type PanelThresholds struct {
	Danger  Threshold `json:"danger"`
	Warning Threshold `json:"warning"`
	Info    Threshold `json:"info"`
}

func NewPanelThresholds() *PanelThresholds {
	d, _ := NewThreshold(0, 0)
	w, _ := NewThreshold(0, 0)
	i, _ := NewThreshold(0, 0)
	return &PanelThresholds{d, w, i}
}
