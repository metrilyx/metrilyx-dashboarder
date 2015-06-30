package datastores

import (
	"encoding/json"
	"fmt"
	"github.com/euforia/simplelog"
	elastigo "github.com/mattbaird/elastigo/lib"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	schemas "github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/schemas/v3"
	"io/ioutil"
	"os"
)

const ESS_DEFAULT_RESULT_SIZE int = 10000000

type EssMapping struct {
	Meta             map[string]interface{} `json:"_meta"`
	DynamicTemplates []interface{}          `json:"dynamic_templates"`
}

type ElasticsearchDatastore struct {
	conn   *elastigo.Conn
	index  string
	logger *simplelog.Logger
}

func NewElasticsearchDatastore(cfg *config.DatastoreConfig, logger *simplelog.Logger) (*ElasticsearchDatastore, error) {

	var (
		ed        = ElasticsearchDatastore{logger: logger}
		err       error
		idxExists bool
	)

	typeCfg, err := ed.parseTypeConfig(cfg.TypeConfig)
	if err != nil {
		return &ed, err
	}

	ed.index = typeCfg.Index

	ed.conn = elastigo.NewConn()
	ed.conn.Domain = typeCfg.Host
	ed.conn.Port = fmt.Sprintf("%d", typeCfg.Port)

	idxExists, err = ed.conn.ExistsIndex(ed.index, "", nil)
	if err != nil {
		if err.Error() == "record not found" {
			idxExists = false
		} else {
			return &ed, err
		}
	}

	if !idxExists {
		return &ed, ed.initializeIndex(typeCfg.MappingFile)
	}
	return &ed, nil
}

func (e *ElasticsearchDatastore) parseTypeConfig(tcfg map[string]interface{}) (*config.DatastoreEssTypeConfig, error) {
	var (
		essCfg = config.DatastoreEssTypeConfig{}
		ok     bool
	)

	if essCfg.MappingFile, ok = tcfg["mapping_file"].(string); !ok {
		return &essCfg, fmt.Errorf("Invalid config param (mapping_file): %s", tcfg)
	}

	if essCfg.Index, ok = tcfg["index"].(string); !ok {
		return &essCfg, fmt.Errorf("Invalid config param (index): %s", tcfg)
	}

	if essCfg.Host, ok = tcfg["host"].(string); !ok {
		return &essCfg, fmt.Errorf("Invalid config param (host): %s", tcfg)
	}
	if val, ok := tcfg["port"].(float64); ok {
		essCfg.Port = int64(val)
	} else {
		return &essCfg, fmt.Errorf("Invalid config param (port): %s", tcfg)
	}

	return &essCfg, nil
}

func (e *ElasticsearchDatastore) initializeIndex(mappingFile string) error {
	resp, err := e.conn.CreateIndex(e.index)
	if err != nil {
		return err
	}
	e.logger.Info.Printf("Index created: %s %s\n", e.index, resp)

	/* apply mapping if specified */
	if mappingFile != "" {
		if _, err := os.Stat(mappingFile); err != nil {
			e.logger.Error.Printf("Mapping file not found %s: %s", mappingFile, err)
			return fmt.Errorf("Mapping file not found %s: %s", mappingFile, err)
		}
		mappingDataBytes, err := ioutil.ReadFile(mappingFile)
		if err != nil {
			return err
		}
		b, err := e.conn.DoCommand("PUT", fmt.Sprintf("/%s/_mapping/_default_", e.index), nil, mappingDataBytes)
		if err != nil {
			return err
		}
		e.logger.Info.Printf("Updated _default_ mapping for %s: %s\n", e.index, b)
	}
	return nil
}

func (e *ElasticsearchDatastore) List() []map[string]interface{} {
	var (
		list    []map[string]interface{}
		tmpDash schemas.Dashboard
	)

	resp, err := e.conn.Search(e.index, "dashboard", nil, "")
	if err != nil {
		return list
	}

	list = make([]map[string]interface{}, len(resp.Hits.Hits))

	for i, hit := range resp.Hits.Hits {

		err := json.Unmarshal(*hit.Source, &tmpDash)
		if err != nil {
			e.logger.Warning.Printf("Could not parse json: %s", *hit.Source)
			list[i] = map[string]interface{}{
				"error": fmt.Sprintf("Could not parse json: %s", *hit.Source),
			}
			continue
		}
		list[i] = map[string]interface{}{
			"name": tmpDash.Name,
			"_id":  tmpDash.Id,
			"tags": tmpDash.Tags,
		}
	}
	return list
}

func (e *ElasticsearchDatastore) Get(_id string) (*schemas.Dashboard, error) {
	var (
		dboard schemas.Dashboard
		err    error
	)

	resp, err := e.conn.Get(e.index, "dashboard", _id, nil)
	if err != nil {
		return &dboard, err
	}

	err = json.Unmarshal(*resp.Source, &dboard)
	if err != nil {
		return &dboard, err
	}
	return &dboard, nil
}

func (e *ElasticsearchDatastore) Add(dboard schemas.Dashboard) error {
	if dboard.Version == 0 {
		dboard.Version = schemas.DASHBOARD_VERSION
	}

	e.logger.Trace.Printf("Adding dashboard: %s\n", dboard.Id)
	_, err := e.conn.Index(e.index, "dashboard", dboard.Id, nil, dboard)
	if err != nil {
		return err
	}

	return nil
}

func (e *ElasticsearchDatastore) Edit(dboard schemas.Dashboard) error {
	e.logger.Trace.Printf("Editing dashboard: %s\n", dboard)
	return e.Add(dboard)
}

func (e *ElasticsearchDatastore) Remove(_id string) error {
	_, err := e.conn.Delete(e.index, "dashboard", _id, nil)
	if err != nil {
		return err
	}

	return nil
}
