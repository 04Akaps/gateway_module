package main

import (
	"flag"
	"github.com/04Akaps/gateway_module/app"
	"github.com/04Akaps/gateway_module/config"
)

var (
	cfg      config.Config
	yamlPath = flag.String("cfgPath", "./deploy.yaml", "config path")
)

func init() {
	flag.Parse()
	cfg = config.NewCfg(*yamlPath)
}

func main() {
	app.NewApp(cfg)
}
