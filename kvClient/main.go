package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type requestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	kvOperation := os.Args[1]
	kvKey := os.Args[2]

	switch {
	case kvOperation == "get":
		getKey(kvKey)
	case kvOperation == "set":
		kvValue := os.Args[3]
		setKey(kvKey, kvValue)
	case kvOperation == "del":
		kvDel(kvKey)
	default:
		log.Printf("Operation '%s' not permitted...", kvOperation)
	}
}

func setKey(key string, value string) {
	httpClient := &http.Client{}

	reqBody := &requestBody{Key: key, Value: value}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("unable to marshal request body: %s", err)
	}

	req := &http.Request{
		Method: "PUT",
		URL: &url.URL{
			Host:   "localhost:8000",
			Path:   "/kv",
			Scheme: "http",
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes)),
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("unable to make put request: %s", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("non 200 HTTP Code from server: %s", resp.Status)
	}

	fmt.Printf("successful response from server for setting key '%s' to value: %s\n", key, value)
}

func kvDel(key string) {
	httpClient := &http.Client{}

	reqBody := &requestBody{Key: key}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("unable to marshal request body: %s", err)
	}

	req := &http.Request{
		Method: "DELETE",
		URL: &url.URL{
			Host:   "localhost:8000",
			Path:   "/kv",
			Scheme: "http",
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes)),
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("unable to make delete request: %s", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("non 200 HTTP Code from server: %s", resp.Status)
	}

	fmt.Printf("key '%s' is successfully deleted\n", key)
}

func getKey(key string) {
	httpClient := &http.Client{}

	reqBody := &requestBody{Key: key}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("unable to marshal request body: %s", err)
	}

	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Host:   "localhost:8000",
			Path:   "/kv",
			Scheme: "http",
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes)),
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("unable to make get request: %s", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("non 200 HTTP Code from server: %s", resp.Status)
	}

	respBody := &requestBody{}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		log.Fatalf("unable to decode response body: %s", err)
	}

	fmt.Printf("value for key '%s' is: %s\n", respBody.Key, respBody.Value)
}
