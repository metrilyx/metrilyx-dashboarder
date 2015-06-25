package services

import (
	"fmt"
	"github.com/euforia/simplelog"
	schemas "github.com/metrilyx/metrilyx-dashboarder/metrilyx/dashboards/schemas/v3"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/providers"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/wrappers"
	"net/http"
	"os"
	"strings"
)

type SchemaHTTPService struct {
	Prefix    string
	Endpoints map[string]string
	logger    *simplelog.Logger
}

func NewSchemaHttpService(prefix, dpath, ppath, gpath, dspath string, logger *simplelog.Logger) *SchemaHTTPService {
	var lgr *simplelog.Logger
	if logger == nil {
		lgr = simplelog.NewLogger(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	} else {
		lgr = logger
	}
	shs := SchemaHTTPService{
		prefix,
		map[string]string{
			dpath:  fmt.Sprintf("%s/%s", prefix, dpath),
			ppath:  fmt.Sprintf("%s/%s", prefix, ppath),
			gpath:  fmt.Sprintf("%s/%s", prefix, gpath),
			dspath: fmt.Sprintf("%s/%s", prefix, dspath),
		},
		lgr,
	}
	shs.registerHandlers()
	return &shs
}

func (s *SchemaHTTPService) registerHandlers() {
	http.HandleFunc(s.Prefix, s.schemaTypesHandler)
	if strings.HasSuffix(s.Prefix, "/") {
		http.HandleFunc(s.Prefix[:len(s.Prefix)-1], s.schemaTypesHandler)
	} else {
		http.HandleFunc(fmt.Sprintf("%s/", s.Prefix), s.schemaTypesHandler)
	}

	s.logger.Warning.Printf("Registering handler: %s", s.Endpoints["dashboard"])
	http.HandleFunc(s.Endpoints["dashboard"], s.dashboardSchemaHandler)
	s.logger.Warning.Printf("Registering handler: %s", s.Endpoints["pod"])
	http.HandleFunc(s.Endpoints["pod"], s.podSchemaHandler)
	s.logger.Warning.Printf("Registering handler: %s", s.Endpoints["panel"])
	http.HandleFunc(s.Endpoints["panel"], s.panelSchemaHandler)

	s.logger.Warning.Printf("Registering handler: %s", s.Endpoints["datasource"])
	http.HandleFunc(s.Endpoints["datasource"], s.datasourceSchemaHandler)
	http.HandleFunc(fmt.Sprintf("%s/", s.Endpoints["datasource"]), s.datasourceSchemaHandler)
}

func (s *SchemaHTTPService) schemaTypesHandler(w http.ResponseWriter, r *http.Request) {
	s.writeJsonReponse(w, r, schemas.SCHEMA_TYPES, 200)
}

func (s *SchemaHTTPService) dashboardSchemaHandler(w http.ResponseWriter, r *http.Request) {
	skeleton := s.isSkeletonRequest(r)
	d := schemas.NewDashboard(skeleton)
	s.writeJsonReponse(w, r, d, 200)
}
func (s *SchemaHTTPService) podSchemaHandler(w http.ResponseWriter, r *http.Request) {
	skeleton := s.isSkeletonRequest(r)
	pod := schemas.NewPod(skeleton)
	s.writeJsonReponse(w, r, pod, 200)
}
func (s *SchemaHTTPService) panelSchemaHandler(w http.ResponseWriter, r *http.Request) {
	//skeleton := s.isSkeletonRequest(r)
	panel := schemas.NewPodPanel()
	s.writeJsonReponse(w, r, panel, 200)
}
func (s *SchemaHTTPService) datasourceSchemaHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if parts[len(parts)-1] == "" || parts[len(parts)-1] == "datasource" {
		/* list provider types */
		s.writeJsonReponse(w, r, providers.PROVIDER_TYPES, 200)
		return
	}

	var (
		providerType = providers.ProviderType(parts[len(parts)-1])
		panelDs      *schemas.PanelDatasource
	)

	panelDs, err := schemas.NewPanelDatasource(providerType)
	if err != nil {
		s.writeJsonReponse(w, r,
			fmt.Sprintf(`{"error": "provider not supported: '%s' %s"}`, providerType, err), 404)
		return
	}
	s.writeJsonReponse(w, r, panelDs, 200)
}

func (s *SchemaHTTPService) isSkeletonRequest(r *http.Request) bool {
	if val, ok := r.URL.Query()["skeleton"]; ok {
		if val[0] == "false" {
			return false
		}
	}
	return true
}

func (s *SchemaHTTPService) writeJsonReponse(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	w.Header().Set("Access-Control-Allow-Origin", ACL_ALLOW_ORIGIN)
	//w.Header().Set("Access-Control-Allow-Headers", ACL_ALLOW_HEADERS)
	wrappers.WriteHttpJsonResponse(w, r, data, code)
	//s.logger.Info.Printf("%d %s\n", code, r.URL.RequestURI())
	s.logger.Info.Printf("%s %d %s\n", r.Method, code, r.URL.RequestURI())
}
