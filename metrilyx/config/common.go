package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

/* Main config file */
type Config struct {
	Datastore DatastoreConfig `json:"datastore"`
	Http      HttpConfig      `json:"http"`
}

type HttpConfig struct {
	Endpoints  HttpEndpointsConfig `toml:"endpoints" json:"endpoints"`
	Port       int                 `toml:"port" json:"port"`
	Webroot    string              `toml:"webroot" json:"webroot"`
	ClientConf string              `toml:"client_conf" json:"client_conf"`
}

type HttpEndpointsConfig struct {
	Dashboards string `toml:"dashboards" json:"dashboards"`
	Schemas    string `toml:"schemas" json:"schemas"`
	Config     string `toml:"config" json:"config"`
}

type DatastoreConfig struct {
	Type       string                 `toml:"type" json:"type"`
	TypeConfig map[string]interface{} `toml:"config" json:"config"`
}

type DatastoreEssTypeConfig struct {
	Host        string `toml:"host" json:"host"`
	Port        int64  `toml:"port" json:"port"`
	Index       string `toml:"index" json:"index"`
	MappingFile string `toml:"mapping_file" json:"mapping_file"`
}

func loadJsonConfig(filepath string) (*Config, error) {
	var (
		cfg Config
		err error
		d   []byte
	)

	d, err = ioutil.ReadFile(filepath)
	if err != nil {
		return &cfg, err
	}

	err = json.Unmarshal(d, &cfg)
	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}

func LoadConfig(cfgfile string) (*Config, error) {
	var (
		cfg *Config
		err error
		wd  string
	)

	if cfg, err = loadJsonConfig(cfgfile); err != nil {
		return cfg, err
	}

	if cfg.Http.Webroot != "" && !strings.HasPrefix(cfg.Http.Webroot, "/") {
		if wd, err = os.Getwd(); err == nil {
			cfg.Http.Webroot = wd + "/" + cfg.Http.Webroot
		} else {
			return cfg, err
		}
	}

	if cfg.Http.Webroot != "" && !strings.HasPrefix(cfg.Http.ClientConf, "/") {
		cfg.Http.ClientConf = cfg.Http.Webroot + "/" + cfg.Http.ClientConf
	}

	return cfg, nil
}
