// +build integration

package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func testHttpGetBody(uri string, output interface{}) (int, error) {
	return internalDo("GET", uri, nil, output)
}

func testHttpGetStatusCode(uri string) (int, error) {
	return internalDo("GET", uri, nil, nil)
}

func testHttpPutModel(uri string, input interface{}, output interface{}) (int, error) {
	return internalDoModel("PUT", uri, input, output)
}

func testHttpPut(uri string, input io.Reader, output interface{}) (int, error) {
	return internalDo("PUT", uri, input, output)
}

func testHttpPostModel(uri string, input interface{}, output interface{}) (int, error) {
	return internalDoModel("POST", uri, input, output)
}

func testHttpPost(uri string, input io.Reader, output interface{}) (int, error) {
	return internalDo("POST", uri, input, output)
}

func testHttpDeleteStatusCode(uri string) (int, error) {
	return internalDo("DELETE", uri, nil, nil)
}

func internalDoModel(method string, uri string, input interface{}, output interface{}) (int, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return 0, err
	}
	return internalDo(method, uri, bytes.NewReader(b), output)
}

func internalDo(method string, uri string, input io.Reader, output interface{}) (int, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("http://web%s", uri), input)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()
	log.Info().Msgf("StatusCode=%d", resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	log.Info().Msgf("Body=%s", string(body))
	if output != nil {
		err = json.Unmarshal(body, output)
		if err != nil {
			return resp.StatusCode, err
		}
	}
	return resp.StatusCode, nil

}
