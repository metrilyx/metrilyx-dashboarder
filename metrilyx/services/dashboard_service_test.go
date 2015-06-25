package services

import (
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	//"net/http"
	"testing"
)

var (
	//testDashId              = "Metrilyx"
	//testDashJsonFile        = "/Users/abs/workbench/GoLang/src/github.com/metrilyx/metrilyx-dashboarder/data/metrilyx.json"
	testConfigFile string = "/Users/abs/workbench/GoLang/src/github.com/metrilyx/metrilyx-dashboarder/etc/metrilyx/metrilyx.json"
	testConfig     *config.Config
	//testEssDs        *ElasticsearchDatastore
	//testDash         = schemas.NewDashboard(true)
	err error
)

func Test_DashboardHTTPService(t *testing.T) {
	if testConfig, err = config.LoadConfig(testConfigFile); err != nil {
		t.Fatalf("%s", err)
	}

	ndhs, err := NewDashboardHTTPService(testConfig, nil)
	if err != nil {
		t.Fatalf("%s", err)
	}

	t.Logf("%v", ndhs)
	//http.ListenAndServe(":6565", nil)
}
