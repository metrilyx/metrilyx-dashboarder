package providers

import (
	"testing"
)

var (
	testHttpProv *HttpProvider
	testUrl      = "http://met3.met.toolsash1.cloudsys.tmcs:443/api/types"
)

func Test_HttpProvider_GET(t *testing.T) {
	testHttpProv = NewHttpProvider()
	testHttpProv.URL = testUrl
	testHttpProv.Method = "GET"
	/*
		_, err := testHttpProv.Fetch()
		if err != nil {
			t.Errorf("%s", err)
		}
	*/
}

func Test_HttpProvider_ErrorMethod(t *testing.T) {
	testHttpProv.Method = "PUT"

	/*_, err := testHttpProv.Fetch()
	if err == nil {
		t.Errorf("Mismatch: %s", err)
	}*/
}
