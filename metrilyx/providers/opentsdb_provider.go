package providers

type OpenTSDBProviderQuery struct {
	Metric       string            `json:"metric"`
	Aggregator   string            `json:"aggregator"`
	Rate         bool              `json:"rate"`
	Tags         map[string]string `json:"tags"`
	MsResolution bool              `json:"msResolution,omitempty"`
	Downsample   string            `json:"downsample,omitempty"`
}

type OpenTSDBRateOptions struct {
	Counter    bool  `json:"counter,omitempty"`
	CounterMax int64 `json:"counterMax,omitempty"`
	ResetValue int64 `json:"resetValue,omitempty"`
}

type OpenTSDBProvider struct {
	HttpProvider

	Query OpenTSDBProviderQuery `json:"query"`
	Start interface{}           `json:"start"`
	End   interface{}           `json:"end,omitempty"`

	RateOptions OpenTSDBRateOptions `json:"rateOptions,omitempty"`
}

func (o *OpenTSDBProvider) DataAggregator() string {
	return o.Query.Aggregator
}
