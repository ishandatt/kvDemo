package httpHelper

import (
	"fmt"
	"kv-for-munish/serverConfig"
	"log"
	"net/http"
)

type kvHandler struct{}

func (kvh kvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Printf("error writing response body: %s", err)
	}
}

// StartHTTP function to listen and serve
func StartHTTP(config serverConfig.ServerConfig) {
	mux := http.NewServeMux()
	mux.Handle("/kv", kvHandler{})

	log.Printf("listening on :%s...", config.ListenPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ListenPort), mux)
	if err != nil {
		log.Fatalf("could not start the http server: %s", err)
	}
}
