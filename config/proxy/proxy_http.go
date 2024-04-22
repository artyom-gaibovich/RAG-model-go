package proxy

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	"rag-model/app/constant"
)

func CreateProxy() *http.Client {
	/*
		Create http.Client with http proxy.
	*/
	proxyURL, err := url.Parse("http://" + constant.GetProxy())
	if err != nil {
		fmt.Println("Ошибка при парсинге адреса прокси:", err)
		return nil
	}

	proxyAuth := &proxy.Auth{
		User:     constant.GetProxyUsername(),
		Password: constant.GetProxyPassword(),
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	if proxyAuth.User != "" && proxyAuth.Password != "" {
		transport.ProxyConnectHeader = http.Header{}
		proxyBasicAuth := proxyAuth.User + ":" + proxyAuth.Password
		proxyAuthHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxyBasicAuth))
		transport.ProxyConnectHeader.Add("Proxy-Authorization", proxyAuthHeader)
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	return httpClient
}
