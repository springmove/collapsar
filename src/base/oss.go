package base

import (
	"fmt"
)

const (
	ServiceOss = "oss"

	Qiniu  = "qiniu"
	Huawei = "huawei"
	S3     = "s3"
)

type IOssService interface {
	Upload(endpoint string, key string, data []byte) error
	Delete(endpoint string, key string) error
}

type Endpoint struct {
	Provider  string `yaml:"provider"`
	AppKey    string `yaml:"app_key"`
	AppSecret string `yaml:"app_secret"`
	Bucket    string `yaml:"bucket"`
	Zone      string `yaml:"zone"`

	// 终端节点, 用于华为云obs
	Endpoint string `yaml:"endpoint"`
}

type IOss interface {
	Init()
	Upload(endpoint string, key string, data []byte) error
	Delete(endpoint string, key string) error
	GetEndpoint(name string) (*Endpoint, error)
	AddEndpoint(name string, endpoint Endpoint)
}

type BaseOss struct {
	Endpoints map[string]Endpoint
}

func (s *BaseOss) Init() {}

func (s *BaseOss) GetEndpoint(name string) (*Endpoint, error) {
	ep, exist := s.Endpoints[name]
	if !exist {
		return nil, fmt.Errorf("Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *BaseOss) AddEndpoint(name string, endpoint Endpoint) {
	if s.Endpoints == nil {
		s.Endpoints = map[string]Endpoint{}
	}
	s.Endpoints[name] = endpoint
}
