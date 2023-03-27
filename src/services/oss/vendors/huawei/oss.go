package huawei

import (
	"bytes"
	"fmt"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/springmove/collapsar/src/base"
	"github.com/springmove/sptty"
)

type Oss struct {
	base.BaseOss

	clients map[string]*obs.ObsClient
}

func (s *Oss) Init() {
	s.clients = map[string]*obs.ObsClient{}
	for name, endpoint := range s.Endpoints {
		client, err := s.initClient(&endpoint)
		if err == nil {
			s.clients[name] = client
		} else {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("huawei.OSS Init Failed: %s", err.Error()))
		}
	}
}

func (s *Oss) Upload(endpoint string, key string, data []byte) error {
	ep, err := s.GetEndpoint(endpoint)
	if err != nil {
		return err
	}

	client, exist := s.clients[endpoint]
	if !exist {
		return fmt.Errorf("Client Not Found ")
	}

	obj := obs.PutObjectInput{}
	obj.Bucket = ep.Bucket
	obj.Key = key
	obj.Body = bytes.NewReader(data)

	_, err = client.PutObject(&obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *Oss) Delete(endpoint string, key string) error {
	ep, err := s.GetEndpoint(endpoint)
	if err != nil {
		return err
	}

	client, exist := s.clients[endpoint]
	if !exist {
		return fmt.Errorf("Client Not Found ")
	}

	obj := obs.DeleteObjectInput{}
	obj.Bucket = ep.Bucket
	obj.Key = key

	_, err = client.DeleteObject(&obj)
	if err != nil {
		return err
	}

	return nil
}

func (s *Oss) initClient(endpoint *base.Endpoint) (*obs.ObsClient, error) {
	obsClient, err := obs.New(endpoint.AppID, endpoint.Secret, endpoint.Endpoint)
	if err != nil {
		return nil, err
	}

	return obsClient, nil
}
