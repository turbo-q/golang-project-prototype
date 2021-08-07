package http

import (
	"golang-project-prototype/config"
	"golang-project-prototype/model"
	"net/http"
	"sync"
	"time"
)

// http client池
// 读多写少，固使用sync.Map
var clientpool sync.Map

func init() {
	client := &http.Client{
		Timeout: time.Duration(config.DefaultConfig.HttpTimeout) * time.Second,
	}
	clientpool.Store(model.CLIENT_DEFAULT, (*httpClient)(client))
}

// 根据flag获取httpclient
// 不同flag使用不同的客户端
// example GetHttpClient("default_client",5)
func GetHttpClient(flag interface{}, timeouts ...int) *httpClient {
	if client, ok := clientpool.Load(flag); ok {
		return client.(*httpClient)
	}

	client := NewHttpClient(timeouts...)
	clientpool.Store(flag, client)
	return client
}

//	创建新的httpclient
func NewHttpClient(timeouts ...int) *httpClient {
	timeout := config.DefaultConfig.HttpTimeout
	if len(timeouts) > 0 {
		timeout = timeouts[0]
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	return (*httpClient)(client)
}
