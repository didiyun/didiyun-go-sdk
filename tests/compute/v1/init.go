package v1

import (
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

var iaasCli *c.DicloudClient

func init() {
	iaasCli = c.InitDicloudClient()
}
