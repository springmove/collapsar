package minio

import (
	"bytes"
	"context"
	"fmt"

	"github.com/linshenqi/collapsar/src/base"
	"github.com/linshenqi/sptty"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Oss struct {
	base.BaseOss

	clients map[string]*minio.Client
}

func (s *Oss) Init() {
	s.clients = map[string]*minio.Client{}

	for name, endpoint := range s.Endpoints {
		client, err := minio.New(endpoint.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(endpoint.AppKey, endpoint.AppSecret, ""),
			Secure: false,
		})

		if err != nil {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("minio.Init.New Failed: %s", err.Error()), "minio")
			continue
		}

		s.clients[name] = client
	}
}

func (s *Oss) getClient(endpoint string) (*minio.Client, *base.Endpoint, error) {
	ep, err := s.GetEndpoint(endpoint)
	if err != nil {
		return nil, nil, err
	}

	client, exist := s.clients[endpoint]
	if !exist {
		return nil, nil, fmt.Errorf("Client Not Found ")
	}

	return client, ep, nil
}

func (s *Oss) Upload(endpoint string, key string, data []byte) error {
	client, ep, err := s.getClient(endpoint)
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = client.PutObject(ctx, ep.Bucket, key, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *Oss) Delete(endpoint string, key string) error {
	client, ep, err := s.getClient(endpoint)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.RemoveObject(ctx, ep.Bucket, key, minio.RemoveObjectOptions{
		ForceDelete: true,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Oss) GetObject(endpoint string, key string) ([]byte, error) {
	client, ep, err := s.getClient(endpoint)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	obj, err := client.GetObject(ctx, ep.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(obj)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
