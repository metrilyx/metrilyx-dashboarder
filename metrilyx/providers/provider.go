package providers

import (
	"encoding/json"
	"fmt"
)

type ProviderType string

const (
	DP_TYPE_OPENTSDB ProviderType = "opentsdb"
	DP_TYPE_HTTP     ProviderType = "http"
)

var PROVIDER_TYPES = []ProviderType{
	DP_TYPE_HTTP,
	DP_TYPE_OPENTSDB,
}

type IProvider interface {
	Fetch() ([]interface{}, error)
}

func NewProvider(ptype ProviderType, pdata []byte) (IProvider, error) {
	var (
		pq  IProvider
		err error
	)

	switch ptype {
	case DP_TYPE_HTTP:
		hpq := &HttpProvider{}
		if err = json.Unmarshal(pdata, hpq); err != nil {
			return pq, err
		}
		if hpq.Method == "" {
			hpq.Method = "GET"
		}
		if hpq.Aggregator == "" {
			hpq.Aggregator = "sum"
		}

		pq = hpq
		break
	case DP_TYPE_OPENTSDB:
		hpq := &OpenTSDBProvider{}
		if err = json.Unmarshal(pdata, hpq); err != nil {
			return pq, err
		}
		if hpq.Method == "" {
			hpq.Method = "GET"
		}
		pq = hpq
		break
	default:
		err = fmt.Errorf("Invalid provider type: %s", ptype)
		break
	}

	return pq, err
}
