package datastores

import (
	"testing"
)

var testDsDir = "/Users/abs/workbench/GoLang/src/github.com/metrilyx/metrilyx-dashboarder/testdata"

func Test_NewJsonFileDatastore(t *testing.T) {
	jds, err := NewJsonFileDatastore(testDsDir)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	l, err := jds.List()
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	t.Logf("%s", l)
}
