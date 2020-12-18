package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/GlebSolncev/http-rest-api/internal/app/apiServer"
	"log"
)
//11:31
var (
	configPath string
)

func init(){
	flag.StringVar(&configPath, "config-path", "configs/apiServer.toml", "Path to config file.")
}

func main()  {
	flag.Parse()
	config := apiServer.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil{
		log.Fatal(err)
	}

	s := apiServer.New(config)
	if err := s.Start(); err != nil{
		log.Fatal(err)
	}


}
