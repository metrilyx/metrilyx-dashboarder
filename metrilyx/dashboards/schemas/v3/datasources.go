package v3

import (
	"encoding/json"
	"fmt"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/providers"
)

type PanelDatasource struct {
	/* Specific data provider ie. http, opentsdb .... */
	Type providers.ProviderType `json:"type"`
	/* provider based on type */
	Provider providers.IProvider `json:"provider"`

	Alias     string `json:"alias"`
	Transform string `json:"transform"`
	Id        string `json:"id"`
}

/*
 * Custom unmarshal to handle different query providers
 * based on provider type
 */
func (p *PanelDatasource) UnmarshalJSON(b []byte) error {
	var (
		tmp map[string]interface{}
		err error

		tVal string
		vOk  bool
	)

	if err = json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	if tVal, vOk = tmp["type"].(string); !vOk {
		return fmt.Errorf("invalid provider type: %s", tmp["type"])
	}
	p.Type = providers.ProviderType(tVal)

	if tVal, vOk = tmp["alias"].(string); vOk {
		p.Alias = tVal
	}

	if tVal, vOk = tmp["transform"].(string); vOk {
		// v3
		p.Transform = tVal
	} else if tVal, vOk = tmp["yTransform"].(string); vOk {
		// v2
		p.Transform = tVal
	}

	if _, vOk := tmp["provider"]; vOk {
		// Marshal query back to bytes to be unmarshaled to it's specific type
		qBytes, _ := json.Marshal(tmp["provider"])
		p.Provider, err = providers.NewProvider(p.Type, qBytes)
	} else {
		// Default http
		p.Provider = providers.NewHttpProvider()
	}

	if p.Id == "" {
		p.Id = NewMetrilyxUUID()
	}
	return err
}

/* TODO - Move this out */
func NewPanelDatasource(ptype providers.ProviderType) (*PanelDatasource, error) {
	var pds = PanelDatasource{}
	//fmt.Println(ptype)
	switch ptype {
	case providers.DP_TYPE_OPENTSDB:
		pds.Provider = &providers.OpenTSDBProvider{
			HttpProvider: providers.HttpProvider{Method: "GET"},
			Query: providers.OpenTSDBProviderQuery{
				Aggregator: "sum",
				Rate:       false,
				Tags:       map[string]string{},
			},
		}
		break
	case providers.DP_TYPE_HTTP:
		pds.Provider = &providers.HttpProvider{Method: "GET"}
		break
	default:
		return &pds, fmt.Errorf("Invalid provider type: %s", ptype)
	}

	pds.Type = ptype
	pds.Id = NewMetrilyxUUID()

	return &pds, nil
}
