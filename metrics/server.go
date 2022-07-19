package metrics

import (
	"fmt"
	"log"
	"net/http"
)

// NewStartedMetricServer create and start an http.Server that listen on `host`:`port`
func NewStartedMetricServer(host string, port int64) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}

	log.Printf("Starting metrics server on %s:%d\n", host, port)
	go log.Fatal(srv.ListenAndServe())

	return &srv
}
