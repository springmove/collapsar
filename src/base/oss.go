package base

import (
	"fmt"
)

const (
	ServiceOss = "oss"

	Qiniu  = "qiniu"
	Huawei = "huawei"
	S3     = "s3"
	MINIO  = "minio"
)

type IOssService interface {
	Upload(endpoint string, key string, data []byte) error
	Delete(endpoint string, key string) error
	ListObjects(endpoint string, prefix string, token string) ([]string, string, error)
	GetObject(endpoint string, key string) ([]byte, error)
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
	ListObjects(endpoint string, prefix string, token string) ([]string, string, error)
	GetEndpoint(name string) (*Endpoint, error)
	AddEndpoint(name string, endpoint Endpoint)
	GetObject(endpoint string, key string) ([]byte, error)
}

type BaseOss struct {
	IOss

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

func (s *BaseOss) ListObjects(endpoint string, prefix string, token string) ([]string, string, error) {
	return nil, "", fmt.Errorf("Not Supported")
}

func (s *BaseOss) GetObject(endpoint string, key string) ([]byte, error) {
	return nil, fmt.Errorf("Not Supported")
}

type ReqBatchDownload struct {
	FileName string
	Key      string
}
