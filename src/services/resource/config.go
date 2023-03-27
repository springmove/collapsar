package resource

import (
	"github.com/springmove/collapsar/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig
	Url         string `yaml:"url"`
	OssEndpoint string `yaml:"ossEndpoint"`
}

func (s *Config) ConfigName() string {
	return base.ServiceResource
}
