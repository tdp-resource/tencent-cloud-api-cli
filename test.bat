@ECHO OFF
::

CD /d %~dp0

CALL go mod tidy

SET "TENCENTCLOUD_REGOIN=ap-guangzhou"
SET "TENCENTCLOUD_SECRET_ID=AKIDxxxxx"
SET "TENCENTCLOUD_SECRET_KEY=xxxxxxxx"

CALL go run main.go ^
    -service "lighthouse" -version "2020-03-24" -action "DescribeInstances" -payload "{}"

IF "%1" == "" CMD /K
