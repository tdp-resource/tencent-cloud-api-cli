package cloud

import (
	"errors"
)

func Proxy(service, version, action, payload, region, secretId, secretKey string) (*Response, error) {

	if service == "" || version == "" || action == "" {
		return nil, errors.New("缺少参数：service，version 或 action")
	}

	if secretId == "" || secretKey == "" {
		return nil, errors.New("缺少参数：secretId 或 secretKey")
	}

	if payload == "" {
		payload = "{}"
	}

	rp := &Params{
		Service:   service,
		Version:   version,
		Action:    action,
		Payload:   []byte(payload),
		Region:    region,
		SecretId:  secretId,
		SecretKey: secretKey,
	}

	return NewRequest(rp)

}
