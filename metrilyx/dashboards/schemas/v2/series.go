package v2

type Serie struct {
	Alias      string `json:"alias"`
	YTransform string `json:"yTransform"`
	Datasource string `json:"datasource"`
	Query      string `json:"query"`
	PaneIndex  int    `json:"paneIndex"`
}

func NewSerie() *Serie {
	return &Serie{"", "", "", "", 0}
}
