package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/GlebSolncev/http-rest-api/app"
	"github.com/GlebSolncev/http-rest-api/internal/apiServer"
)

var (
	configPath string
)

// constructor ...
func init() {
	flag.StringVar(&configPath, "config", "config.toml", "Path to config file.")
}

func main() {
	flag.Parse()
	config := apiServer.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	app.Check(err)

	s := apiServer.New(config)
	err = s.Start()
	app.Check(err)
}
