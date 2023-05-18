package oss

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/springmove/collapsar/src/base"
)

var endpoint string = "test"

func getService() *Service {
	srv := Service{cfg: Config{Endpoints: map[string]base.Endpoint{
		endpoint: {
			Provider: base.MINIO,
			AppID:    "",
			Secret:   "",
			Bucket:   "",
			Endpoint: "",
			SSL:      false,
		},
	}}}

	srv.setupProviders()
	srv.initProviders()

	return &srv
}

func getUrlImage(url string) ([]byte, *http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, &resp.Header, nil
}

func TestService(t *testing.T) {
	srv := getService()

	// content, err := ioutil.ReadFile("/Users/linshenqi/temp/Johnrogershousemay2020.webp")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// if err := srv.Upload("asset", filename, content); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	filename := "test"

	data, header, err := getUrlImage("https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTI7YP85icEWiaYmNaYFUrvNdjiaCcFMxO0Wj9dH7QfX5ZtibHPW50HDsibLoHApUic5EALnja3vnn0uKziaQ/132")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := srv.Upload(endpoint, filename, data, minio.PutObjectOptions{
		ContentType: header.Get("content-type"),
	}); err != nil {
		fmt.Println(err.Error())
		return
	}

	// if err := srv.Delete(endpoint, filename); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

}
