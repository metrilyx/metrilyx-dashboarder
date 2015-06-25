package datastores

import (
	"github.com/euforia/simplelog"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"testing"
)

func Test_LoadDatastore(t *testing.T) {
	testConfig, _ = config.LoadConfig(testConfigFile)
	ds, err := LoadDatastore(&testConfig.Datastore, simplelog.NewStdLogger())
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%v", ds)
}

func Test_LoadDatastore_NOT_SUPPORTED(t *testing.T) {
	testConfig.Datastore.Type = "hoopla"
	if _, err := LoadDatastore(&testConfig.Datastore, simplelog.NewStdLogger()); err == nil {
		t.Fatalf("%s", err)
	}
}
