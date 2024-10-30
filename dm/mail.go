package dm

import (
	"net/http"

	"github.com/hiifong/aliyungo/common"
)

type SendEmailArgs struct {
	AccountName string
	AddressType string
}

type SendBatchMailArgs struct {
	SendEmailArgs
	TemplateName string
	ReceiverName string
	TagName      string
}

//remember to setup the accountName in your aliyun console
//addressType should be "1" or "0",
//0:random address, it's recommanded
//1:sender's address
//tagName is optional, you can use "" if you don't wanna use it
//please set the receiverName and template in the console of Aliyun before you call this API,if you use tagName, you should set it as well

func (this *Client) SendBatchMail(args *SendBatchMailArgs) error {
	return this.InvokeByAnyMethod(http.MethodPost, BatchSendMail, "", args, &common.Response{})
}

type SendSingleMailArgs struct {
	SendEmailArgs
	ReplyToAddress bool
	ToAddress      string
	FromAlias      string
	Subject        string
	HtmlBody       string
	TextBody       string
}

//remember to setup the accountName in your aliyun console
//addressType should be "1" or "0",
//0:random address, it's recommanded
//1:sender's address
//please set the receiverName and template in the console of Aliyun before you call this API,if you use tagName, you should set it as well

//fromAlias, subject, htmlBody, textBody are optional

func (this *Client) SendSingleMail(args *SendSingleMailArgs) error {
	return this.InvokeByAnyMethod(http.MethodPost, SingleSendMail, "", args, &common.Response{})
}

// https://help.aliyun.com/zh/direct-mail/delivery-overview?spm=a2c4g.11186623.help-menu-29412.d_8_10.1ecd453dIYbvlg

type SenderStatisticsDetailByParamArgs struct {
	AccountName string `json:"AccountName,omitempty"` // 发新地址
	TagName     string `json:"TagName,omitempty"`     // 邮件标签
	StartTime   string `json:"StartTime,omitempty"`   // 起始时间
	EndTime     string `json:"EndTime,omitempty"`     // 结束时间
	ToAddress   string `json:"ToAddress,omitempty"`   // 收心地址
	Status      int    `json:"Status,omitempty"`      // 投递结果
	Length      int    `json:"Length,omitempty"`      // 请求返回的结果数目，1-100
	NextStart   string `json:"NextStart,omitempty"`   // 分页用
}

type SenderStatisticsDetailByParamResponse struct {
	common.Response
	Data struct {
		MailDetail []MailDetail `json:"mailDetail"`
	} `json:"data"`
	NextStart string
}

type MailDetail struct {
	ToAddress           string
	LastUpdateTime      string
	UtcLastUpdateTime   int
	AccountName         string
	Message             string
	Status              int
	Subject             string
	ErrorClassification string
}

func (this *Client) SenderStatisticsDetailByParam(args *SenderStatisticsDetailByParamArgs) (resp *SenderStatisticsDetailByParamResponse, err error) {
	resp = &SenderStatisticsDetailByParamResponse{}
	err = this.InvokeByAnyMethod(http.MethodPost, SenderStatisticsDetailByParam, "", args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
