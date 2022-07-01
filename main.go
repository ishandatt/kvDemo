package main

import (
	"kv-for-munish/httpHelper"
	sc "kv-for-munish/serverConfig"
	"log"
	"os"
)

func main() {
	var config sc.ServerConfig
	sc.ReadConfigFile(&config)

	createKVDataDir(config)

	httpHelper.StartHTTP(config)
}

func createKVDataDir(config sc.ServerConfig) {
	err := os.MkdirAll(config.KvPath, 0755)
	if err != nil {
		log.Fatalf("could not create the kv data directory: %s", err)
	}
}
