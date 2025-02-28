package core

import (
	"net/http"
	"net/url"
	"time"
)

func NewHttpClient(timeout int, proxy string) *http.Client {
	transport := &http.Transport{}
	if proxy != "" {
		proxyUrl, _ := url.Parse(proxy)
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: transport,
	}
	return client
}