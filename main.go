package main

import (
	"fmt"
	sc "kv-for-munish/serverConfig"
)

func main() {
	var config sc.ServerConfig
	sc.ReadConfigFile(&config)
	fmt.Println(config)
}
