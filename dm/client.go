package dm

import (
	"github.com/hiifong/aliyungo/common"
)

const (
	EmailEndPoint                 = "https://dm.aliyuncs.com/"
	SingleSendMail                = "SingleSendMail"
	BatchSendMail                 = "BatchSendMail"
	SenderStatisticsDetailByParam = "SenderStatisticsDetailByParam"
	EmailAPIVersion               = "2015-11-23"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := new(Client)
	client.Init(EmailEndPoint, EmailAPIVersion, accessKeyId, accessKeySecret)
	return client
}
