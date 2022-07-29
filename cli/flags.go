package cli

import (
	"flag"
	"fmt"
	"os"
)

var (
	service string
	version string
	action  string
	payload string

	region    string
	secretId  string
	secretKey string
)

func flags() {

	flag.StringVar(&service, "service", "", "服务名\n")
	flag.StringVar(&version, "version", "", "服务版本\n")
	flag.StringVar(&action, "action", "", "服务接口名\n")
	flag.StringVar(&payload, "payload", "", "接口入参，默认值为 {}\n")

	flag.StringVar(&region, "region", "", "地域，默认值为环境变量 TENCENTCLOUD_REGOIN\n")
	flag.StringVar(&secretId, "secretId", "", "密钥 ID，默认值为环境变量 TENCENTCLOUD_SECRET_ID\n")
	flag.StringVar(&secretKey, "secretKey", "", "密钥 KEY，默认值为环境变量 TENCENTCLOUD_SECRET_KEY\n")

	flag.Usage = usage

	flag.Parse()

}

func usage() {

	fmt.Fprintf(os.Stderr, `使用方法:

  export TENCENTCLOUD_REGOIN="ap-guangzhou"
  export TENCENTCLOUD_SECRET_ID="AKIDxxxxx"
  export TENCENTCLOUD_SECRET_KEY="xxxxxxxx"

  tcapi -service "cvm" -version "2017-03-12" -action "DescribeRegions" -payload "{}"

  tcapi -service "cvm" -version "2017-03-12" -action "DescribeRegions" -payload "{}" -region "ap-guangzhou" -secretId "xx" -secretKey "xxx"

参数列表:

`)

	flag.PrintDefaults()

}
