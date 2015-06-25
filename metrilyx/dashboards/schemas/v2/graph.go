package v2

type Graph struct {
	Chart
	MultiPane   bool             `json:"multiPane"`
	Panes       []string         `json:"panes"`
	Thresholds  *ChartThresholds `json:"threshold"`
	Series      []*Serie         `json:"series"`
	Secondaries []*Serie         `json:"secondaries"`
}

func NewGraph(skeleton bool) *Graph {
	nc := NewChart()
	if skeleton {
		return &Graph{
			Chart:       *nc,
			MultiPane:   false,
			Panes:       []string{"", ""},
			Thresholds:  NewChartThresholds(),
			Series:      make([]*Serie, 0),
			Secondaries: make([]*Serie, 0),
		}
	} else {
		return &Graph{
			Chart:       *nc,
			MultiPane:   false,
			Panes:       []string{"", ""},
			Thresholds:  NewChartThresholds(),
			Series:      []*Serie{NewSerie()},
			Secondaries: make([]*Serie, 0),
		}
	}
}
