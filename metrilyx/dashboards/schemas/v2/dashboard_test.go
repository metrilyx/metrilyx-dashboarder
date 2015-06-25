package v2

import (
	"encoding/json"
	"testing"
)

//var testDashboardFile string = ""

func Test_Dashboard(t *testing.T) {
	nd := NewDashboard(true)
	if len(nd.Layout[0][0]) != 0 {
		t.Errorf("length mismatch: %s", nd.Layout[0][0])
		t.FailNow()
	}

	nd = NewDashboard(false)
	if nd.Layout[0][0][0].Graphs[0].Type != "line" {
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

/*
func Test_NewDashboardFromFile(t *testing.T) {
	d, err := NewDashboardFromFile(testDashboardFile)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	t.Logf("%s", d)
}
*/
