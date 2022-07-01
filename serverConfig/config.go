package serverConfig

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

// ServerConfig to hold config
type ServerConfig struct {
	KvPath     string `yaml:"kv_path"`
	ListenPort string `yaml:"listen_port"`
}

// ReadConfigFile to read server config
func ReadConfigFile(config *ServerConfig) {
	kvServerConfigFile := getConfigFilePath()
	fileContents, err := os.ReadFile(kvServerConfigFile)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	err = yaml.Unmarshal(fileContents, &config)
	if err != nil {
		log.Fatalf("unable to unmarshal config yaml: %s", err)
	}
}

func getConfigFilePath() string {
	var kvServerConfigFile string
	if len(os.Args) > 1 {
		if strings.Split(os.Args[1], "=")[0] != "-c" {
			log.Printf("unrecognised command line option '%s'", strings.Split(os.Args[1], "=")[0])
			log.Fatalf("please provide config path using '%s'", "-c=")
		} else {
			kvServerConfigFile = strings.Split(os.Args[1], "=")[1]
			log.Printf("starting server with config file at %s", kvServerConfigFile)
		}
	} else {
		log.Fatalf("please provide config path using '%s'", "-c=")
	}
	return kvServerConfigFile
}
