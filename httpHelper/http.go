package httpHelper

import (
	"fmt"
	"kv-for-munish/serverConfig"
	"log"
	"net/http"
)

func (kvh kvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		kvh.getKV(w, r)
	case r.Method == "PUT":
		kvh.setKV(w, r)
	case r.Method == "DELETE":
		kvh.deleteKV(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// StartHTTP function to listen and serve
func StartHTTP(config serverConfig.ServerConfig) {
	mux := http.NewServeMux()
	mux.Handle("/kv", kvHandler{
		kvPath: config.KvPath,
	})

	log.Printf("listening on :%s...", config.ListenPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ListenPort), mux)
	if err != nil {
		log.Fatalf("could not start the http server: %s", err)
	}
}
