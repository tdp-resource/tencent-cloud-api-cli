package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"tcapi/cloud"
)

func Exec() {

	flags()

	if rg, ok := os.LookupEnv("TENCENTCLOUD_REGOIN"); ok && region == "" {
		region = rg
	}

	if ak, ok := os.LookupEnv("TENCENTCLOUD_SECRET_ID"); ok && secretId == "" {
		secretId = ak
	}

	if sk, ok := os.LookupEnv("TENCENTCLOUD_SECRET_KEY"); ok && secretKey == "" {
		secretKey = sk
	}

	res, err := cloud.Proxy(service, version, action, payload, region, secretId, secretKey)

	if err == nil {
		json, _ := json.Marshal(res.Response)
		fmt.Println(string(json))
	} else {
		fmt.Println(err)
	}

}
