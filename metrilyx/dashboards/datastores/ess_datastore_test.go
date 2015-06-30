package datastores

import (
	"encoding/json"
	"github.com/euforia/simplelog"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	schemas "github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/schemas/v3"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var (
	testLogger          = simplelog.NewStdLogger()
	testDashId          = "Metrilyx"
	testDashJsonFile, _ = filepath.Abs("../../../testdata/metrilyx.json")
	testConfigFile, _   = filepath.Abs("../../../etc/metrilyx/metrilyx.json.sample")
	testEssMapFile, _   = filepath.Abs("../../../etc/metrilyx/ess-dashboard-mapping.json")
	testConfig          *config.Config
	testEssDs           *ElasticsearchDatastore
	testDash            = schemas.NewDashboard(true)
)

func Test_ElasticsearchDatastore(t *testing.T) {
	var (
		err error
	)

	testConfig, err = config.LoadConfig(testConfigFile)
	if err != nil {
		t.Fatalf("FAILED: %s", err)
	}
	testConfig.Datastore.TypeConfig["mapping_file"] = testEssMapFile

	testEssDs, err = NewElasticsearchDatastore(&testConfig.Datastore, testLogger)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
}

func Test_ElasticsearchDatastore_NOT_FOUND(t *testing.T) {
	if _, err := testEssDs.Get("hoopla-1232"); err == nil {
		t.Errorf("Error checking failed")
	}
}

func Test_ElasticsearchDatastore_Add(t *testing.T) {
	/* Read test dashboard json */
	b, err := ioutil.ReadFile(testDashJsonFile)
	if err = json.Unmarshal(b, &testDash); err != nil {
		t.Fatalf("%s", err)
	}

	if err = testEssDs.Add(*testDash); err != nil {
		t.Fatalf("Failed to add: %s", err)
	}
}

func Test_ElasticsearchDatastore_Get_Edit(t *testing.T) {

	gotDash, err := testEssDs.Get(testDashId)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if err = testEssDs.Edit(*testDash); err != nil {
		t.Fatalf("%s", err)
	}

	//b, err = json.MarshalIndent(gotDash, "", "  ")
	t.Logf("%s\n", gotDash.Id)
}

func Test_ElasticsearchDatastore_List(t *testing.T) {

	list := testEssDs.List()
	/*
	 * This will always be empty due to the time take by ess to index,
	 * as this is a _search call.
	 */
	t.Logf("%#v\n", list)
}

func Test_ElasticsearchDatastore_Remove(t *testing.T) {
	if err := testEssDs.Remove(testDashId); err != nil {
		t.Errorf("%s", err)
	}
	if err := testEssDs.Remove(testDashId); err == nil {
		t.Errorf("error check failed")
	}
}
