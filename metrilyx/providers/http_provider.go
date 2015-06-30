package providers

type HttpProvider struct {
	/* Used when multiple of the same type */
	Name   string              `json:"name"`
	URL    string              `json:"url"`
	Method string              `json:"method"`
	Params map[string][]string `json:"params,omitempty"`
	/* Http body */
	Body interface{} `json:"body,omitempty"`
	/* Normalize response data to fit pandas.DataFrame dict */
	Normalizer string `json:"normalizer,omitempty"`

	Aggregator string `json:"aggregator"`
}

func NewHttpProvider() *HttpProvider {
	return &HttpProvider{
		Method:     "GET",
		Aggregator: DEFAULT_AGGR,
	}
}

func (h *HttpProvider) DataAggregator() string {
	return h.Aggregator
}
