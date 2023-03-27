package main

import (
	"flag"

	"github.com/springmove/collapsar/src/services/oss"
	"github.com/springmove/collapsar/src/services/resource"
	"github.com/springmove/sptty"
)

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := sptty.GetApp()
	app.ConfFromFile(*cfg)

	services := sptty.Services{
		&oss.Service{},
		&resource.Service{},
	}

	configs := sptty.Configs{
		&oss.Config{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
