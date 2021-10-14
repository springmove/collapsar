package oss

import (
	"fmt"
	"testing"

	"github.com/linshenqi/collapsar/src/base"
)

func getService() *Service {
	srv := Service{cfg: Config{Endpoints: map[string]base.Endpoint{
		"asset": {
			Provider:  base.MINIO,
			AppKey:    "admin",
			AppSecret: "",
			Bucket:    "asset",
			Endpoint:  "godev.xiji.space:39000",
		},
	}}}

	srv.setupProviders()
	srv.initProviders()

	return &srv
}

func TestService(t *testing.T) {
	srv := getService()

	filename := "f1.webp"
	// content, err := ioutil.ReadFile("/Users/linshenqi/temp/Johnrogershousemay2020.webp")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// if err := srv.Upload("asset", filename, content); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// if err := srv.Delete("asset", filename); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	f, err := srv.GetObject("asset", filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(len(f))

}
