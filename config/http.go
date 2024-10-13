package config

import "github.com/04Akaps/gateway_module/types/http"

type HttpCfg struct {
	BaseURL string `yaml:"base_url"`
	Router  router `yaml:"router"`
	Path    string `yaml:"path"`

	Auth   *Auth             `yaml:"auth"`
	Header map[string]string `yaml:"header"`
}

type router struct {
	Method   http.HttpMethod `yaml:"method"`
	GetType  http.GetType    `yaml:"get_type"`
	Variable []string        `yaml:"variable"`
	//UrlVariable []string        `yaml:"url_variable"`
}

type Auth struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
