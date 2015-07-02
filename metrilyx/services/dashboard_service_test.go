package services

import (
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"path/filepath"
	"testing"
)

var (
	testEssMapFile, _ = filepath.Abs("../../../etc/metrilyx/ess-dashboard-mapping.json")
)

func Test_DashboardHTTPService(t *testing.T) {
	var err error
	if testConfig, err = config.LoadConfig(testConfigFile); err != nil {
		t.Fatalf("%s", err)
	}

	testConfig.Datastore.TypeConfig["mapping_file"] = testEssMapFile

	ndhs, err := NewDashboardHTTPService(testConfig, nil)
	if err != nil {
		t.Fatalf("%s", err)
	}

	t.Logf("%v", ndhs)
}
