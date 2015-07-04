package services

import (
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

var (
	testConfigFile, _   = filepath.Abs("../../etc/metrilyx/metrilyx.json.sample")
	testUIConfigFile, _ = filepath.Abs("../../etc/metrilyx/ui.json.sample")
	testConfig          *config.Config
)

func Test_ConfigService(t *testing.T) {

	cfg, err := config.LoadConfig(testConfigFile)
	if err != nil {
		t.Fatalf("%s", err)
	}
	cfg.Http.ClientConf = testUIConfigFile
	t.Logf("Endpoint: %s", cfg.Http.Endpoints.Config)

	go func() {

		if _, err = NewConfigService(cfg, nil); err != nil {
			t.Fatalf("%s", err)
		}
		http.ListenAndServe(":3456", nil)
	}()
	time.Sleep(3)
	resp, err := http.Get("http://localhost:3456/api/config")
	if err != nil {
		t.Fatalf("%s", err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Wrong status code: %d %s", resp.StatusCode, b)
	}

	t.Logf("%s", b)

}
