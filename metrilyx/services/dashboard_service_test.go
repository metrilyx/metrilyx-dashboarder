package services

import (
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"testing"
)

func Test_DashboardHTTPService(t *testing.T) {
	var err error
	if testConfig, err = config.LoadConfig(testConfigFile); err != nil {
		t.Fatalf("%s", err)
	}

	ndhs, err := NewDashboardHTTPService(testConfig, nil)
	if err != nil {
		t.Fatalf("%s", err)
	}

	t.Logf("%v", ndhs)
}
