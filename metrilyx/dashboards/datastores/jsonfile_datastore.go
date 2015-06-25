package datastores

import (
	"encoding/json"
	schemas "github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/schemas/v3"
	"io/ioutil"
	"os"
)

/*
	IN PROGRESS
*/
type JsonFileDatastore struct {
	DashboardsDir string
}

func NewJsonFileDatastore(dirpath string) (*JsonFileDatastore, error) {
	j := JsonFileDatastore{}
	if _, err := os.Stat(dirpath); err != nil {
		return &j, err
	}
	j.DashboardsDir = dirpath
	return &j, nil
}

func (j *JsonFileDatastore) List() ([]map[string]string, error) {
	filelist, err := ioutil.ReadDir(j.DashboardsDir)
	if err != nil {
		return make([]map[string]string, 0), err
	}
	namelist := make([]map[string]string, len(filelist))
	for i, f := range filelist {
		if !f.IsDir() {
			namelist[i] = map[string]string{"name": f.Name()}
		}
	}
	return namelist, nil
}

func (j *JsonFileDatastore) Get(id string) (*schemas.Dashboard, error) {
	var (
		d   schemas.Dashboard
		err error
		b   []byte
	)

	if b, err = ioutil.ReadFile(j.DashboardsDir + id + ".json"); err == nil {
		if err = json.Unmarshal(b, &d); err == nil {
			return &d, nil
		}
	}
	return &d, err
}

/*
func (j *JsonFileDatastore) Add(schemas.Dashboard) error                       {}
func (j *JsonFileDatastore) Remove(id string) error                            {}
func (j *JsonFileDatastore) Edit(schemas.Dashboard) (schemas.Dashboard, error) {}
*/
