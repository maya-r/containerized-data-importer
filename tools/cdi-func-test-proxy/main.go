package main

import (
	"github.com/elazarl/goproxy"
	"sync"
	"log"
	"net/http"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})
	go func() {
		log.Fatal(http.ListenAndServe(":8080", proxy))
		wg.Done()
	}()

	proxyAuth := goproxy.NewProxyHttpServer()
	proxyAuth.Verbose = true
	proxyAuth.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			user, pass, _ := r.BasicAuth()
			if "aaa" != user || "bbb" != pass {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusUnauthorized,
					"Basic auth wrong")
			}
			r.Header.Set("X-GoProxy", "yxorPoG-X")
			return r, nil
		})

	go func() {
		log.Fatal(http.ListenAndServe(":8081", proxy))
		wg.Done()
	}()

	wg.Wait()
}
