package cloud

import (
	"errors"
)

func Request(service, version, action, payload, region, secretId, secretKey string) (*TakeResponse, error) {

	if service == "" || version == "" || action == "" {
		return nil, errors.New("缺少参数：service，version 或 action")
	}

	if secretId == "" || secretKey == "" {
		return nil, errors.New("缺少参数：secretId 或 secretKey")
	}

	if payload == "" {
		payload = "{}"
	}

	c := NewClient(region, secretId, secretKey)

	return c.Take(service, version, action, payload)

}
