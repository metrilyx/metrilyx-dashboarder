package services

import (
	"encoding/json"
	"github.com/euforia/simplelog"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/wrappers"
	"io/ioutil"
	"net/http"
)

/*
 * Client (UI) config endpoint service .
 */

type ConfigService struct {
	ConfigFile string
	/* Endpoints from config to serve to client */
	epConfig config.HttpEndpointsConfig

	logger *simplelog.Logger
}

func NewConfigService(cfg *config.Config, logger *simplelog.Logger) (*ConfigService, error) {
	cfgSvc := ConfigService{
		ConfigFile: cfg.Http.ClientConf,
		epConfig:   cfg.Http.Endpoints,
	}

	if logger != nil {
		cfgSvc.logger = logger
	} else {
		cfgSvc.logger = simplelog.NewStdLogger()
	}

	_, err := ioutil.ReadFile(cfgSvc.ConfigFile)
	if err != nil {
		return &cfgSvc, err
	}

	cfgSvc.logger.Info.Printf("UI config: %s\n", cfgSvc.ConfigFile)
	cfgSvc.RegisterHandler()

	return &cfgSvc, nil
}

func (c *ConfigService) RegisterHandler() {
	c.logger.Info.Printf("Registering config handler: %s\n", c.epConfig.Config)
	http.HandleFunc(c.epConfig.Config, c.configHandler)
}

func (c *ConfigService) configHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var tmp = map[string]interface{}{
			"endpoints": c.epConfig,
		}

		b, err := ioutil.ReadFile(c.ConfigFile)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		err = json.Unmarshal(b, &tmp)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		wrappers.WriteHttpJsonResponse(w, r, tmp, 200)
		break
	/*case "POST":
		break
	case "PUT":
		break
	case "PATCH":
		break*/
	default:
		w.WriteHeader(405)
		break
	}
}
