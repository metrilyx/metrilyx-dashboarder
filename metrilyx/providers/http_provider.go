package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/metrilyx/metrilyx-dashboarder/metrilyx/wrappers"
)

const DEFAULT_AGGR = "sum"

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

	httpClient *wrappers.HTTPCall
}

func NewHttpProvider() *HttpProvider {
	return &HttpProvider{
		httpClient: wrappers.NewHTTPCall(false),
		Method:     "GET",
		Aggregator: DEFAULT_AGGR,
	}
}

func (h *HttpProvider) URLEncodedParams() string {
	if h.Params == nil || len(h.Params) < 1 {
		return ""
	}

	out := "?"
	for k, vals := range h.Params {
		for _, v := range vals {
			out += k + "=" + v + "&"
		}
	}

	return out[:len(out)-1]
}

func (h *HttpProvider) requestBody() ([]byte, error) {
	return json.Marshal(h.Body)
}

func (h *HttpProvider) Fetch() ([]interface{}, error) {
	var (
		httpResp wrappers.HttpResponseData
		err      error
		url      = h.URL + h.URLEncodedParams()
		resp     []interface{}
	)

	if h.URL == "" {
		return resp, fmt.Errorf("URL missing!")
	}
	// In case NewHttpProvider was not used
	if h.httpClient == nil {
		h.httpClient = wrappers.NewHTTPCall(false)
	}

	switch h.Method {
	case "GET":
		httpResp, err = h.httpClient.Get(url)
		break
	case "POST":
		data, err := h.requestBody()
		if err != nil {
			break
		}
		httpResp, err = h.httpClient.Post(url, "application/json", bytes.NewBuffer(data))
		break
	default:
		err = fmt.Errorf("Method not supported: %s", h.Method)
		break
	}

	if err != nil {
		return resp, err
	}

	err = httpResp.AsJson(&resp)

	return resp, err
}
