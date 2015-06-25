package v3

/*
 * Pod contents. Can be a graph, text etc...
 */
import (
	"code.google.com/p/go-uuid/uuid"
	"strings"
)

type PodPanelSize string
type PodPanelType string

const (
	PANEL_SIZE_SMALL  PodPanelSize = "small"
	PANEL_SIZE_MEDIUM PodPanelSize = "medium"
	PANEL_SIZE_LARGE  PodPanelSize = "large"
)

const DEFAULT_PANEL_SIZE PodPanelSize = PANEL_SIZE_MEDIUM

const DEFAULT_PANEL_REFRESH = 50

const (
	PANEL_TYPE_AREA    PodPanelType = "area"
	PANEL_TYPE_BAR     PodPanelType = "bar"
	PANEL_TYPE_COLUMN  PodPanelType = "column"
	PANEL_TYPE_LINE    PodPanelType = "line"
	PANEL_TYPE_PIE     PodPanelType = "pie"
	PANEL_TYPE_STACKED PodPanelType = "stacked"

	PANEL_TYPE_IMAGE   PodPanelType = "image"
	PANEL_TYPE_GRAPHIC PodPanelType = "graphic"
	PANEL_TYPE_LIST    PodPanelType = "list"
	PANEL_TYPE_TEXT    PodPanelType = "text"
)

const (
	GRAPH_LIB_HIGHCHARTS = "highcharts"
	GRAPH_LIB_DYGRAPHS   = "dygraphs"
)

const (
	DEFAULT_PANEL_TYPE PodPanelType = PANEL_TYPE_LINE
	DEFAULT_GRAPH_LIB               = GRAPH_LIB_HIGHCHARTS
)

func NewMetrilyxUUID() string {
	return strings.Join(strings.Split(uuid.New(), "-"), "")
}

type Secondary struct {
	Operation  string `json:"operation"`
	Alias      string `json:"alias"`
	Aggregator string `json:"aggregator"`
}

type PodPanel struct {
	Id   string       `json:"id"`
	Name string       `json:"name"`
	Size PodPanelSize `json:"size"`
	Type PodPanelType `json:"type"`

	/* JS graphing library */
	GraphingLibrary string `json:"library"`
	/*
	 * Options passed to graphics library.
	 * These are exposed to the ui for user options.
	 */
	Graphics map[string]interface{} `json:"graphics"`
	/* Refresh interval */
	Refresh int64 `json:"refresh"`
	/*
	 * Datasources used to populate the panel.
	 * Order needs to be preserved.
	 * e.g. multiple tsdb queries, ess query etc.
	 */
	Datasources []PanelDatasource `json:"datasources"`
	Thresholds  *PanelThresholds  `json:"thresholds"`

	Secondaries []Secondary `json:"secondaries"`
}

/*
 * Args:
 *  Id
 *  Name
 */
func NewPodPanel(args ...string) *PodPanel {
	panelId := NewMetrilyxUUID()

	if len(args) > 0 && args[0] != "" {
		panelId = args[0]
	}

	cThresh := NewPanelThresholds()

	switch len(args) {
	case 2:
		return &PodPanel{
			Id:              panelId,
			Name:            args[1],
			Size:            DEFAULT_PANEL_SIZE,
			Type:            DEFAULT_PANEL_TYPE,
			Graphics:        map[string]interface{}{},
			GraphingLibrary: DEFAULT_GRAPH_LIB,
			Refresh:         DEFAULT_PANEL_REFRESH,
			Datasources:     make([]PanelDatasource, 0),
			Thresholds:      cThresh,
			Secondaries:     make([]Secondary, 0),
		}
	default:
		return &PodPanel{
			Id:              panelId,
			Name:            "",
			Size:            DEFAULT_PANEL_SIZE,
			Type:            DEFAULT_PANEL_TYPE,
			Graphics:        map[string]interface{}{},
			GraphingLibrary: DEFAULT_GRAPH_LIB,
			Refresh:         DEFAULT_PANEL_REFRESH,
			Datasources:     make([]PanelDatasource, 0),
			Thresholds:      cThresh,
			Secondaries:     make([]Secondary, 0),
		}
	}
}
