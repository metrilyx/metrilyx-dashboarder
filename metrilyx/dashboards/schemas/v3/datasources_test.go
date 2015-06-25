package v3

import (
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/providers"
	"testing"
)

func Test_NewPanelDatasource_HTTP(t *testing.T) {
	if _, err := NewPanelDatasource(providers.DP_TYPE_HTTP); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}

}

func Test_NewPanelDatasource_OPENTSDB(t *testing.T) {
	if _, err := NewPanelDatasource(providers.DP_TYPE_OPENTSDB); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
}

func Test_NewPanelDatasource_INVALID(t *testing.T) {
	if _, err := NewPanelDatasource(providers.ProviderType("hoopla")); err == nil {
		t.Errorf("Error check failed")
	}
}
