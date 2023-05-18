package oss

import (
	"fmt"

	"github.com/springmove/collapsar/src/base"
	"github.com/springmove/collapsar/src/services/oss/vendors/huawei"
	"github.com/springmove/collapsar/src/services/oss/vendors/minio"
	"github.com/springmove/collapsar/src/services/oss/vendors/qiniu"
	"github.com/springmove/collapsar/src/services/oss/vendors/s3"
	"github.com/springmove/sptty"
)

type Service struct {
	sptty.BaseService
	base.IServiceOss

	cfg       Config
	providers map[string]base.IOss
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	s.setupProviders()
	s.initProviders()

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceOss
}

func (s *Service) initProviders() {
	for k, v := range s.cfg.Endpoints {
		provider, err := s.getProvider(v.Provider)
		if err != nil {
			continue
		}

		provider.AddEndpoint(k, v)
	}

	for _, provider := range s.providers {
		provider.Init()
	}
}

func (s *Service) getProvider(providerType string) (base.IOss, error) {
	provider, exist := s.providers[providerType]
	if !exist {
		return nil, fmt.Errorf("Provider Not Found ")
	}

	return provider, nil
}

func (s *Service) getEndpoint(endpoint string) (*base.Endpoint, error) {
	ep, exist := s.cfg.Endpoints[endpoint]
	if !exist {
		return nil, fmt.Errorf("Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *Service) setupProviders() {
	s.providers = map[string]base.IOss{
		base.Qiniu:  &qiniu.Oss{},
		base.Huawei: &huawei.Oss{},
		base.S3:     &s3.Oss{},
		base.MINIO:  &minio.Oss{},
	}
}

func (s *Service) Upload(endpoint string, key string, data []byte, opt ...interface{}) error {
	ep, err := s.getEndpoint(endpoint)
	if err != nil {
		return err
	}

	provider, err := s.getProvider(ep.Provider)
	if err != nil {
		return err
	}

	return provider.Upload(endpoint, key, data, opt...)
}

func (s *Service) UploadFromFile(endpoint string, key string, filepath string, opt ...interface{}) error {
	ep, err := s.getEndpoint(endpoint)
	if err != nil {
		return err
	}

	provider, err := s.getProvider(ep.Provider)
	if err != nil {
		return err
	}

	return provider.UploadFromFile(endpoint, key, filepath, opt...)
}

func (s *Service) Delete(endpoint string, key string) error {
	ep, err := s.getEndpoint(endpoint)
	if err != nil {
		return err
	}

	provider, err := s.getProvider(ep.Provider)
	if err != nil {
		return err
	}

	return provider.Delete(endpoint, key)
}

func (s *Service) ListObjects(endpoint string, prefix string, token string) ([]string, string, error) {
	ep, err := s.getEndpoint(endpoint)
	if err != nil {
		return nil, "", err
	}

	provider, err := s.getProvider(ep.Provider)
	if err != nil {
		return nil, "", err
	}

	return provider.ListObjects(endpoint, prefix, token)
}

func (s *Service) GetObject(endpoint string, key string) ([]byte, error) {
	ep, err := s.getEndpoint(endpoint)
	if err != nil {
		return nil, err
	}

	provider, err := s.getProvider(ep.Provider)
	if err != nil {
		return nil, err
	}

	return provider.GetObject(endpoint, key)
}

func (s *Service) GetConfig() map[string]base.Endpoint {
	return s.cfg.Endpoints
}
