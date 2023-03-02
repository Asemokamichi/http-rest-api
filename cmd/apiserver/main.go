package main

import (
	"flag"
	"log"

	"github.com/Asemokamichi/http-rest-api/internal/app/apiserver"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
