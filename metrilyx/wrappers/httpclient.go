package wrappers

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpResponseData struct {
	bytes   []byte
	Headers map[string][]string
}

func NewHttpResponseData(dataBytes []byte) HttpResponseData {
	return HttpResponseData{bytes: dataBytes}
}

func (h *HttpResponseData) AsJson(iface interface{}) error {
	return json.Unmarshal(h.bytes, iface)
}

func (h *HttpResponseData) AsString() string {
	return string(h.bytes)
}

func (h *HttpResponseData) Bytes() []byte {
	return h.bytes
}

type HTTPCall struct {
	client http.Client
}

func NewHTTPCall(secure bool) *HTTPCall {
	httpTransport := &http.Transport{
		DisableCompression: false,
	}
	if secure {
		httpTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return &HTTPCall{client: http.Client{Transport: httpTransport}}
}

func (h *HTTPCall) getResponse(r *http.Response, e error) (HttpResponseData, error) {
	//fmt.Printf("%v\n", r.Header)
	var (
		hrd = HttpResponseData{Headers: r.Header}
		err error
	)

	if e != nil {
		return hrd, e
	}
	defer r.Body.Close()

	hrd.bytes, err = ioutil.ReadAll(r.Body)

	return hrd, err
}

func (h *HTTPCall) Get(urlStr string) (HttpResponseData, error) {
	return h.getResponse(h.client.Get(urlStr))
}

func (h *HTTPCall) Post(url string, datatype string, data io.Reader) (HttpResponseData, error) {

	return h.getResponse(h.client.Post(url, datatype, data))
}

func WriteHttpJsonResponse(w http.ResponseWriter, r *http.Request, data interface{}, respCode int) {
	var b []byte
	s, ok := data.(string)
	if ok {
		b = []byte(s)
	} else {
		if _, ok := r.URL.Query()["pretty"]; ok {
			b, _ = json.MarshalIndent(&data, "", "  ")
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		} else {
			b, _ = json.Marshal(&data)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
		}
	}

	w.WriteHeader(respCode)
	w.Write(b)
}
