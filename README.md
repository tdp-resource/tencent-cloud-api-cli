# 腾讯云跨平台命令行工具

## 使用方法:

```shell
  export TENCENTCLOUD_REGOIN="ap-guangzhou"
  export TENCENTCLOUD_SECRET_ID="AKIDxxxxx"
  export TENCENTCLOUD_SECRET_KEY="xxxxxxxx"

  tcapi -service "cvm" -version "2017-03-12" -action "DescribeRegions" -payload "{}"

  tcapi -service "cvm" -version "2017-03-12" -action "DescribeRegions" -payload "{}" -region "ap-guangzhou" -secretId "xx" -secretKey "xxx"
```

## 参数列表:

```
  -action string
        服务接口名

  -payload string
        接口入参，默认值为 {}

  -region string
        地域，默认值为环境变量 TENCENTCLOUD_REGOIN

  -secretId string
        密钥 ID，默认值为环境变量 TENCENTCLOUD_SECRET_ID

  -secretKey string
        密钥 KEY，默认值为环境变量 TENCENTCLOUD_SECRET_KEY

  -service string
        服务名

  -version string
        服务版本
```
