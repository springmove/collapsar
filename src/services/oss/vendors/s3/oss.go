package s3

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/linshenqi/collapsar/src/base"
	"github.com/linshenqi/sptty"
)

type Oss struct {
	base.BaseOss

	clients map[string]*s3.S3
}

func (s *Oss) Init() {
	s.clients = map[string]*s3.S3{}
	for name, endpoint := range s.Endpoints {
		session, err := session.NewSession(&aws.Config{
			Region:           aws.String(endpoint.Zone),
			Credentials:      credentials.NewStaticCredentials(endpoint.AppID, endpoint.Secret, ""),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(false),
		})

		if err != nil {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("s3.Init.NewSession Failed: %s", err.Error()), "S3")
			continue
		}

		s.clients[name] = s3.New(session)
	}
}

func (s *Oss) getClient(endpoint string) (*s3.S3, *base.Endpoint, error) {
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

	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(ep.Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})

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

	_, err = client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(ep.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Oss) ListObjects(endpoint string, prefix string, token string) ([]string, string, error) {
	client, ep, err := s.getClient(endpoint)
	if err != nil {
		return nil, "", err
	}

	req := s3.ListObjectsV2Input{
		Bucket: aws.String(ep.Bucket),
		Prefix: aws.String(prefix),
	}

	if token != "" {
		req.ContinuationToken = aws.String(token)
	}

	output, err := client.ListObjectsV2(&req)

	if err != nil {
		return nil, "", err
	}

	rt := []string{}
	for _, v := range output.Contents {
		if *v.Size > 0 {
			rt = append(rt, fmt.Sprintf("%s/%s/%s", ep.Endpoint, ep.Bucket, *v.Key))
		}
	}

	newToken := ""
	if output.NextContinuationToken != nil {
		newToken = *output.NextContinuationToken
	}

	return rt, newToken, nil
}

// func (s *Oss) BatchUploadFromFiles(endpoint string, key string, data []byte) error {
// 	client, _, err := s.getClient(endpoint)
// 	if err != nil {
// 		return err
// 	}

// 	client.getobj
// 	return nil
// }

func (s *Oss) BatchDownloadToFile(endpoint string, req []*base.ReqBatchDownload) error {

	return nil
}
