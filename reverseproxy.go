package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func reverseProxy(mux *http.ServeMux, config map[string]string) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uriPath := r.URL.Path

		//-- handle landing of main
		if uriPath == "/" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("This is main app"))
			return
		}

		//-- handle reverse proxy
		for pattern, target := range config {
			if strings.HasPrefix(uriPath, pattern) {
				reverseProxyHandler(pattern, target)(w, r)
				return
			}
		}

		//-- no handle
		http.NotFound(w, r)
	})
}

func reverseProxyHandler(pattern, target string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		appURL, err := url.Parse(target)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s, target: %s", err.Error(), target), http.StatusInternalServerError)
			return
		}

		requestPath := r.URL.Path
		if strings.HasPrefix(requestPath, pattern) && requestPath != pattern {
			requestPath = requestPath[len(pattern):]
			r.URL.Path = requestPath
		}
		w.Header().Set("X-ReverseProxy", "MyAwesomeReverseProxy")

		proxy := httputil.NewSingleHostReverseProxy(appURL)
		proxy.ServeHTTP(w, r)
	}
}
