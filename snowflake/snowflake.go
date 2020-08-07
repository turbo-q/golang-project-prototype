package snowflake

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var cnf *snowflakeConfig

func init() {
	snowflakeConfigInit()
	cnf = &snowflakeConfig{
		Domain:     SnowflakeConfig.Domain,
		AuthUser:   SnowflakeConfig.AuthUser,
		AuthSecret: SnowflakeConfig.AuthSecret,
	}
}

var curlIDClient *http.Client

type curlRespID struct {
	FID string `json:"F_id"`
}

func init() {
	curlIDClient = &http.Client{
		Transport: &http.Transport{
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 5,
		},
	}
}

// GetUniqueID 获取唯一ID
func GetUniqueID() (id string) {
	uri := cnf.Domain + "/v1/snowflak/id"
	method := "GET"
	req, _ := http.NewRequest(method, uri, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authentication", cnf.AuthUser+":"+md5Sum(cnf.AuthSecret))

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

func md5Sum(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}
