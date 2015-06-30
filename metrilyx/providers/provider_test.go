package providers

import (
	"encoding/json"
	"testing"
)

var (
	testHttpProvData, _ = json.Marshal(map[string]string{
		"url":        "testurl",
		"aggregator": "min",
	})
	testTsdbProvData, _ = json.Marshal(map[string]interface{}{
		"url":   "testurl",
		"query": map[string]string{"aggregator": "min"},
	})
)

func Test_NewHttpProvider(t *testing.T) {
	prov := NewHttpProvider()
	if prov.Method != "GET" || prov.Aggregator != DEFAULT_AGGR {
		t.Fatalf("Mismatch!")
	}
}

func Test_NewProvider_OpenTSDB(t *testing.T) {

	prov, err := NewProvider(DP_TYPE_OPENTSDB, testTsdbProvData)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if prov.DataAggregator() != "min" {
		t.Fatalf("Mismatch: %s", prov.DataAggregator())
	}
}

func Test_NewProvider_Http(t *testing.T) {

	prov, err := NewProvider(DP_TYPE_HTTP, testHttpProvData)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if prov.DataAggregator() != "min" {
		t.Fatalf("Mismatch: %s", prov.DataAggregator())
	}
}
