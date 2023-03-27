package oss

import "github.com/springmove/collapsar/src/base"

type Config struct {
	Endpoints map[string]base.Endpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return base.ServiceOss
}

func (s *Config) Validate() error {
	return nil
}

func (s *Config) Default() interface{} {
	return &Config{}
}
