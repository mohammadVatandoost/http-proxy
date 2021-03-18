package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
	"fmt"
)

func main() {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
	fmt.Println("Proxy run in port: 7575")
	proxy.OnRequest().DoFunc(
	func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
		fmt.Println("Proxy request")
		return r,nil
	})
    log.Fatal(http.ListenAndServe(":7575", proxy))
	fmt.Println("Proxy runing")
}

