package httpHelper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

type requestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type kvHandler struct {
	kvPath string
}

func (kvh kvHandler) getKV(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("received get request for key '%s'", reqBody.Key)

	keyPath := fmt.Sprintf("%s/%s", kvh.kvPath, reqBody.Key)

	keyValue, err := os.ReadFile(keyPath)
	if err != nil {
		log.Printf("error reading key: %s", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	reqBody.Value = string(keyValue)
	err = json.NewEncoder(w).Encode(reqBody)
	if err != nil {
		log.Printf("error writing response body: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (kvh kvHandler) setKV(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("received set request for key '%s' with value '%s'", reqBody.Key, reqBody.Value)

	keyPath := fmt.Sprintf("%s/%s", kvh.kvPath, reqBody.Key)

	err = os.MkdirAll(path.Dir(keyPath), 0755)
	if err != nil {
		log.Printf("error creating directory for key: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = os.WriteFile(keyPath, []byte(reqBody.Value), 0744)
	if err != nil {
		log.Printf("error persisting key: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (kvh kvHandler) deleteKV(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("received delete request for key '%s'", reqBody.Key)

	keyPath := fmt.Sprintf("%s/%s", kvh.kvPath, reqBody.Key)

	err = os.Remove(keyPath)
	if err != nil {
		log.Printf("error deleting key: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
