package main

import (
	"flag"
	"fmt"
	"github.com/euforia/simplelog"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/services"
	"net/http"
	"os"
)

var (
	configfile = flag.String("c", "metrilyx.json", "Path to configuration file")
	loglevel   = flag.String("l", "info", "Log level [ trace | debug | info | warning | error ]")
)

func main() {
	flag.Parse()

	var (
		cfg    *config.Config
		logger = simplelog.NewLogger(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
		err    error
	)

	logger.SetLogLevel(*loglevel)

	if cfg, err = config.LoadConfig(*configfile); err != nil {
		logger.Error.Fatalf("%s\n", err)
	}

	/* Schema service */
	services.NewSchemaHttpService(cfg.Http.Endpoints.Schemas,
		"dashboard", "pod", "panel", "datasource", logger)

	/* Dashboard service */
	if _, err = services.NewDashboardHTTPService(cfg, logger); err != nil {
		logger.Error.Printf("%s\n", err)
	}

	/* Config handler */
	if _, err = services.NewConfigService(cfg, logger); err != nil {
		logger.Error.Fatalf("Config: %s; Reason: %s\n", cfg.Http.ClientConf, err)
	}

	/* Webroot handler */
	http.Handle("/", http.FileServer(http.Dir(cfg.Http.Webroot)))
	logger.Warning.Printf("Webroot: %s\n", cfg.Http.Webroot)

	logger.Warning.Printf("Starting HTTP server on port %d...\n", cfg.Http.Port)
	logger.Error.Fatalf("%s", http.ListenAndServe(fmt.Sprintf(":%d", cfg.Http.Port), nil))
}
