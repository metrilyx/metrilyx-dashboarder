package config

import (
	"path/filepath"
	"testing"
)

var (
	testConfigFile, _ = filepath.Abs("../../etc/metrilyx/metrilyx.json")
	testErrCfgFile, _ = filepath.Abs("../../testdata/bad_json.json")
)

func Test_LoadConfig(t *testing.T) {

	testConfig, err := LoadConfig(testConfigFile)
	if err != nil {
		t.Errorf("FAILED: %s", err)
		t.FailNow()
	}
	t.Logf("%#v", testConfig)
}

func Test_LoadConfig_ErrorNotFound(t *testing.T) {
	_, err := LoadConfig("/not/existing/path")
	if err == nil {
		t.Fatalf("Error check failed: file existence")
	}
}

func Test_LoadConfig_ErrorBadJson(t *testing.T) {
	_, err := LoadConfig(testErrCfgFile)
	if err == nil {
		t.Fatalf("Error check failed: load error")
	}
}
