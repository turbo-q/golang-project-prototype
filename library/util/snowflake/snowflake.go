package snowflake

import (
	"encoding/json"
	"golang-project-prototype/config"
	"golang-project-prototype/library/helper"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var curlIDClient *http.Client

type curlRespID struct {
	FID string `json:"F_id"`
}

type curlRespIntID struct {
	FID int `json:"F_id"`
}

func init() {
	curlIDClient = &http.Client{
		Transport: &http.Transport{
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 5,
		},
	}
}

// GetStringID 获取唯一stringId（长度20位）
func GetStringId() (id string) {
	uri := config.SnowflakeConfig.Domain + "/v1/snowflak/id"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authentication", config.SnowflakeConfig.AuthUser+":"+helper.MD5Sum(config.SnowflakeConfig.AuthSecret))

	// logs.FormRequest(appCtx, remoteName, req, nil)

	idObj := curlRespID{}
	if resp, err := curlIDClient.Do(req); err == nil {
		defer resp.Body.Close()
		if respBody, err := ioutil.ReadAll(resp.Body); err == nil {
			dec := json.NewDecoder(strings.NewReader(string(respBody)))
			dec.UseNumber()
			dec.Decode(&idObj)
		}
		// logs.Response(appCtx, remoteName, req, resp, respBody, time.Now().Sub(start))

		if resp.Status == "200 OK" {
			id = idObj.FID
		}
	} else {
		// logs.ResponseErr(appCtx, remoteName, req, err, time.Now().Sub(start))
	}

	return
}

// GetStringID 获取唯一intId（长度16位）
func GetUniqueIntId() (id int) {
	id = 0

	uri := config.SnowflakeConfig.Domain + "/v1/snowflak/intId"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authentication", config.SnowflakeConfig.AuthUser+":"+helper.MD5Sum(config.SnowflakeConfig.AuthSecret))

	resp, err := curlIDClient.Do(req)
	idObj := curlRespIntID{}
	if err == nil {
		defer resp.Body.Close()
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
			dec.UseNumber()
			dec.Decode(&idObj)
		}
		if resp.Status == "200 OK" {
			id = idObj.FID
		}
		//log response
		// logapp.LogCurlResponse("向发号器'获取int类型id'发起请求响应", string(bodyByte), resp.Status, appCtx)
	} else {
		//log err
		// logapp.LogCurlResponseErr("向发号器'获取int类型id'发起请求响应报错:", err, appCtx)
	}
	return
}
