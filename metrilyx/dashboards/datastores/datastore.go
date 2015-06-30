package datastores

import (
	"fmt"
	"github.com/euforia/simplelog"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	schemas "github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/schemas/v3"
)

type IDashboardDatastore interface {
	List() []map[string]interface{}
	Get(string) (*schemas.Dashboard, error)
	Add(schemas.Dashboard) error
	Remove(string) error
	Edit(schemas.Dashboard) error
}

func LoadDatastore(cfg *config.DatastoreConfig, logger *simplelog.Logger) (IDashboardDatastore, error) {
	switch cfg.Type {
	case "elasticsearch":
		return NewElasticsearchDatastore(cfg, logger)
	default:
		return nil, fmt.Errorf("datastore type not supported: %s", cfg.Type)
	}
}
