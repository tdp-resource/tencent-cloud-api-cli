package main

import (
	"encoding/json"
	"fmt"

	"tcapi/cloud"
)

func main() {

	res, err := cloud.Request(
		"lighthouse", "2020-03-24", "DescribeInstances", "{}",
		"ap-guangzhou", "xxx", "xxx",
	)

	if err == nil {
		json, _ := json.Marshal(res.Response)
		fmt.Println(string(json))
	} else {
		fmt.Println(err)
	}

}
