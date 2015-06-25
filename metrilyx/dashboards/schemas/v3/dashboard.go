package v3

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

const DASHBOARD_VERSION int = 3

type Dashboard struct {
	Name        string     `json:"name"`
	Id          string     `json:"_id"`
	Tags        []string   `json:"tags"`
	Layout      [][][]*Pod `json:"layout"`
	Version     int        `json:"version"`
	LastUpdated int64      `json:"_lastUpdated"`
}

func NewDashboard(skeleton bool) *Dashboard {
	if skeleton {
		return &Dashboard{
			Name:        "",
			Id:          "",
			Tags:        make([]string, 0),
			Layout:      [][][]*Pod{{{}}},
			Version:     DASHBOARD_VERSION,
			LastUpdated: time.Now().UnixNano(),
		}
	} else {
		return &Dashboard{
			Name:        "",
			Id:          "",
			Tags:        make([]string, 0),
			Layout:      [][][]*Pod{{{NewPod(false)}}},
			Version:     DASHBOARD_VERSION,
			LastUpdated: time.Now().UnixNano(),
		}
	}
}

func NewDashboardFromFile(dpath string) (*Dashboard, error) {
	var d Dashboard
	b, err := ioutil.ReadFile(dpath)
	if err != nil {
		return &d, err
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return &d, err
	}
	return &d, nil
}
