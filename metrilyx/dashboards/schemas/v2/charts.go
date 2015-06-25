package v2

/*
 * Pod contents. Can be a graph, text etc...
 */
import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"strings"
)

type ChartSize string

const (
	CHART_SIZE_SMALL  ChartSize = "small"
	CHART_SIZE_MEDIUM ChartSize = "medium"
	CHART_SIZE_LARGE  ChartSize = "large"
)

const DEFAULT_CHART_SIZE ChartSize = CHART_SIZE_MEDIUM

type ChartType string

const (
	CHART_TYPE_AREA    ChartType = "area"
	CHART_TYPE_BAR     ChartType = "bar"
	CHART_TYPE_COLUMN  ChartType = "column"
	CHART_TYPE_LINE    ChartType = "line"
	CHART_TYPE_PIE     ChartType = "pie"
	CHART_TYPE_STACKED ChartType = "stacked"
)

const DEFAULT_CHART_TYPE ChartType = CHART_TYPE_LINE

type Chart struct {
	Id   string    `json:"_id"`
	Name string    `json:"name"`
	Size ChartSize `json:"size"`
	Type ChartType `json:"type"`
}

/*
 * Args:
 *  id
 *  name
 */
func NewChart(args ...string) *Chart {
	chartId := strings.Join(strings.Split(uuid.New(), "-"), "")

	if len(args) > 0 && args[0] != "" {
		chartId = args[0]
	}
	switch len(args) {
	case 1:
		return &Chart{chartId, "", DEFAULT_CHART_SIZE, DEFAULT_CHART_TYPE}
	case 2:
		return &Chart{chartId, args[1], DEFAULT_CHART_SIZE, DEFAULT_CHART_TYPE}
	default:
		return &Chart{chartId, "", DEFAULT_CHART_SIZE, DEFAULT_CHART_TYPE}
	}
}

type Threshold struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

func NewThreshold(max, min float64) (Threshold, error) {
	var t Threshold
	if max < min {
		return t, fmt.Errorf("max cannot be less than min")
	}
	t = Threshold{max, min}
	return t, nil
}

type ChartThresholds struct {
	Danger  Threshold `json:"danger"`
	Warning Threshold `json:"warning"`
	Info    Threshold `json:"info"`
}

func NewChartThresholds() *ChartThresholds {
	d, _ := NewThreshold(0, 0)
	w, _ := NewThreshold(0, 0)
	i, _ := NewThreshold(0, 0)
	return &ChartThresholds{d, w, i}
}
