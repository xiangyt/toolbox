package proxy

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestNewProxy(t *testing.T) {
	host := os.Getenv("proxy")
	if host == "" {
		host = "http://nas.zerotier.xyt:5700/"
	}
	// initialize a reverse proxy and pass the actual backend server url here
	// 初始化反向代理并传入真正后端服务的地址
	proxy, err := NewProxy(host)
	if err != nil {
		panic(err)
	}

	// handle all requests to your server using the proxy
	// 使用 proxy 处理所有请求到你的服务
	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(":9320", nil))
}
