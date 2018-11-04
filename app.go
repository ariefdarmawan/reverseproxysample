package main

import "net/http"

func helloApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, greeting from HelloApp"))
	})

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + name))
	})

	hs := &http.Server{Addr: ":9123", Handler: mux}
	_ = hs.ListenAndServe()
}

func welcomeApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, greeting from WelcomeApp"))
	})

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome " + name))
	})

	hs := &http.Server{Addr: ":9124", Handler: mux}
	_ = hs.ListenAndServe()
}

func proxyApp() {
	mux := http.NewServeMux()
	reverseProxy(mux, map[string]string{
		"/hello":   "http://localhost:9123",
		"/welcome": "http://localhost:9124",
	})

	hs := &http.Server{Addr: ":9090", Handler: mux}
	_ = hs.ListenAndServe()
}
