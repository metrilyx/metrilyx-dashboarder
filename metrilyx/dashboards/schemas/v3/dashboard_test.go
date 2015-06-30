package v3

import (
	"encoding/json"
	"testing"
)

func Test_Dashboard(t *testing.T) {
	nd := NewDashboard(true)
	if len(nd.Layout[0][0]) != 0 {
		t.Errorf("length mismatch: %s", nd.Layout[0][0])
		t.FailNow()
	}

	nd = NewDashboard(false)
	if nd.Layout[0][0][0].Panels[0].Type != "line" {
		t.Errorf("graph type mismatch: %s", nd)
		t.FailNow()
	}

	b, err := json.MarshalIndent(&nd, "", "  ")
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	t.Logf("%s\n", b)
}
