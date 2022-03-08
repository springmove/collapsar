package resource

import (
	"github.com/linshenqi/collapsar/src/base"
	"github.com/linshenqi/sptty"
)

type Config struct {
	sptty.BaseConfig
	Url         string `yaml:"url"`
	OssEndpoint string `yaml:"ossEndpoint"`
}

func (s *Config) ConfigName() string {
	return base.ServiceResource
}
