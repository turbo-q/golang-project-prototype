package http

import (
	"encoding/json"
	"golang-project-prototype/library/util/logger"
	"golang-project-prototype/model"
	"io/ioutil"
	"net/http"
	"net/url"
)

type httpClient http.Client

func (client httpClient) GetByReceiver(u string, values url.Values, receiver interface{}) error {
	request, _ := http.NewRequest(http.MethodGet, u+"?"+values.Encode(), nil)
	cp := http.Client(client)
	resp, err := cp.Do(request)
	if err != nil {
		logger.Error("http 请求出错", err)
		return model.ErrRequest
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, receiver)
	if err != nil {
		return model.ErrJSONDecode
	}

	return nil
}
