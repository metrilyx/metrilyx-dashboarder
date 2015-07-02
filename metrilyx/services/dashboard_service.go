package services

import (
	"encoding/json"
	"fmt"
	"github.com/euforia/simplelog"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/config"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/datastores"
	schemas "github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/schemas/v3"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/wrappers"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ACL_ALLOW_ORIGIN  string = "*"
	ACL_ALLOW_HEADERS string = "content-type, authorization, accept"
	ACL_ALLOW_METHODS string = "GET, POST, PUT, OPTIONS, DELETE"
)

type DashboardHTTPService struct {
	datastore     datastores.IDashboardDatastore
	endpointParts []string
	logger        *simplelog.Logger
}

func NewDashboardHTTPService(cfg *config.Config, logger *simplelog.Logger) (*DashboardHTTPService, error) {
	var (
		d = DashboardHTTPService{
			logger:        logger,
			endpointParts: make([]string, 0),
		}
		err error
	)

	if d.logger == nil {
		d.logger = simplelog.NewStdLogger()
	}

	d.logger.Trace.Printf("Loading datastore with config: %v...\n", cfg.Datastore)
	if d.datastore, err = datastores.LoadDatastore(&cfg.Datastore, d.logger); err != nil {
		return &d, err
	}
	d.logger.Debug.Printf("Datastore loaded: %s", cfg.Datastore.Type)

	d.setEndpointParts(cfg.Http.Endpoints.Dashboards)
	d.registerEndpoints(cfg.Http)

	return &d, nil
}

func (d *DashboardHTTPService) listHandler(w http.ResponseWriter, r *http.Request) {

	var (
		resp interface{}
		code int = 200
	)
	switch r.Method {
	case "OPTIONS":
		resp = map[string]string{"message": "todo: Print help for endpoint"}
		w.Header().Set("Access-Control-Allow-Origin", ACL_ALLOW_ORIGIN)
		break
	case "GET":
		resp = d.datastore.List()
		break
	default:
		resp = fmt.Sprintf(`{"error": "Method not supported: %s"}`, r.Method)
		code = 501
		break

	}
	d.writeJsonReponse(w, r, resp, code)
}

func (d *DashboardHTTPService) editorHandler(w http.ResponseWriter, r *http.Request) {

	var (
		pathParts   = d.parseURLPath(r.URL.Path)
		dashboardId string
	)

	if len(pathParts) == 0 {
		d.listHandler(w, r)
		return
	}

	if len(pathParts) > 1 {
		d.writeJsonReponse(w, r, map[string]string{
			"error": fmt.Sprintf("Invalid request %s %s", r.Method, r.URL.Path)}, http.StatusNotAcceptable)
		return
	}

	dashboardId = pathParts[0]

	var (
		resp   interface{}
		code   int = 200
		err    error
		dboard *schemas.Dashboard
	)

	switch r.Method {
	case "OPTIONS":
		resp = map[string]string{"message": "todo: Print help for endpoint"}
		w.Header().Set("Access-Control-Allow-Headers", ACL_ALLOW_HEADERS)
		w.Header().Set("Access-Control-Allow-Methods", ACL_ALLOW_METHODS)
		break
	case "GET":
		resp, err = d.datastore.Get(dashboardId)
		break
	case "POST":
		dboard, err = d.readBodyAsDashboard(dashboardId, r)
		if err != nil {
			d.logger.Warning.Printf("Failed to add dashboard: %s\n", err)
		} else {
			err = d.datastore.Add(*dboard)
			resp = `{"status": "success"}`
		}
		break
	case "PUT":
		dboard, err = d.readBodyAsDashboard(dashboardId, r)
		if err != nil {
			d.logger.Warning.Printf("Failed to edit dashboard: %s\n", err)
		} else {
			err = d.datastore.Edit(*dboard)
			resp = `{"status": "success"}`
		}
		break
	case "DELETE":
		err = d.datastore.Remove(dashboardId)
		resp = `{"status": "success"}`
		break
	default:
		d.writeJsonReponse(w, r,
			fmt.Sprintf(`{"error": "Method not supported: %s"}`, r.Method), 501)
		return
	}

	if err != nil {
		resp = map[string]string{"error": err.Error()}

		if err.Error() == "record not found" {
			code = 404
		} else if code >= 200 && code <= 304 {
			code = 500
		}
	}

	d.writeJsonReponse(w, r, resp, code)
}

func (d *DashboardHTTPService) registerEndpoints(cfg config.HttpConfig) {

	if strings.HasSuffix(cfg.Endpoints.Dashboards, "/") {
		http.HandleFunc(cfg.Endpoints.Dashboards[:len(cfg.Endpoints.Dashboards)-1], d.listHandler)
		http.HandleFunc(cfg.Endpoints.Dashboards, d.editorHandler)
	} else {
		http.HandleFunc(cfg.Endpoints.Dashboards, d.listHandler)
		http.HandleFunc(fmt.Sprintf("%s/", cfg.Endpoints.Dashboards), d.editorHandler)
	}
}

func (d *DashboardHTTPService) setEndpointParts(endpoint string) {
	for _, v := range strings.Split(endpoint, "/") {
		if v == "" {
			continue
		}
		d.endpointParts = append(d.endpointParts, v)
	}
}

func (d *DashboardHTTPService) parseURLPath(urlpath string) []string {
	pathParts := strings.Split(urlpath, "/")
	out := make([]string, 0)
	for _, v := range pathParts {
		if v == "" {
			continue
		}
		out = append(out, v)
	}
	return out[len(d.endpointParts):]
}

func (d *DashboardHTTPService) readBodyAsDashboard(id string, r *http.Request) (*schemas.Dashboard, error) {
	var dboard schemas.Dashboard
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &dboard, err
	}
	err = json.Unmarshal(body, &dboard)
	if err != nil {
		return &dboard, err
	}
	dboard.Id = id
	return &dboard, nil
}

func (s *DashboardHTTPService) writeJsonReponse(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	w.Header().Set("Access-Control-Allow-Origin", ACL_ALLOW_ORIGIN)
	wrappers.WriteHttpJsonResponse(w, r, data, code)
	s.logger.Info.Printf("%s %d %s\n", r.Method, code, r.URL.RequestURI())
}
