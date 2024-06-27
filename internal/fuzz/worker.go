package fuzz

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/kmarkela/duffman/internal/pcollection"
)

func startWorker(wg *sync.WaitGroup, wq <-chan workUnit, wr chan<- workResults, tr *http.Transport) {

	for wu := range wq {

		if !wu.parBody {

			getParam := make(map[string]string)
			for k, v := range wu.r.Parameters.Get {
				getParam[k] = v
			}
			getParam[wu.param] = wu.word

			endpoint := createEndpoint(wu.r.URL, getParam)
			var r io.Reader = strings.NewReader(wu.r.Body)
			doRequest(endpoint, r, wu, wr, tr)

			continue
		}

		body, _ := encodeBody(wu)
		endpoint := createEndpoint(wu.r.URL, wu.r.Parameters.Get)
		doRequest(endpoint, body, wu, wr, tr)
	}

	wg.Done()
}

func createEndpoint(url string, par map[string]string) string {
	endpoint := fmt.Sprintf("%s?", url)
	for gk, gv := range par {
		endpoint = fmt.Sprintf("%s%s=%s&", endpoint, gk, gv)
	}

	return endpoint
}

func doRequest(endpoint string, body io.Reader, wu workUnit, wr chan<- workResults, tr *http.Transport) {

	result := workResults{
		endpoint: wu.r.URL,
		param:    wu.param,
		word:     wu.word,
	}

	req, err := http.NewRequest(wu.r.Method, endpoint, body)
	if err != nil {
		result.
		return err
	}

	for k, v := range wu.r.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func encodeBody(wu workUnit) (io.Reader, error) {

	// to restore param back
	postParam := make(map[string]string)
	for k, v := range wu.r.Parameters.Post {
		postParam[k] = v
	}

	postParam[wu.param] = wu.word

	// encode Form
	if strings.HasPrefix(wu.r.ContentType, "application/x-www-form-urlencoded") {
		form := url.Values{}
		for k, v := range postParam {
			form.Add(k, v)
		}
		return strings.NewReader(form.Encode()), nil
	}

	// encode json
	if wu.r.ContentType == "application/json" {
		b := pcollection.MarshalJSONBody(postParam)
		return bytes.NewBuffer(b), nil
	}

	// TODO: encode multipart

	// unknown content type
	return strings.NewReader(wu.r.Body), fmt.Errorf("no encoder for: %s", wu.r.ContentType)
}
