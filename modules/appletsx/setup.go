package appletsx

import (
	"gin_template/modules/component"

	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/payment"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

var applets client

func init() {
	component.RegComponent(&applets)
}

const (
	ConfSectionWxApplets = "wxapplets"
)

type client struct {
	miniProgram      *miniProgram.MiniProgram `ini:"-"`
	payment          *payment.Payment         `ini:"-"`
	AppID            string
	AppSecret        string
	MsgCallbackToken string
	MchID            string
	MchApiV2Key      string
	MchApiV3Key      string
	CertPath         string
	KeyPath          string
	SerialNo         string
	NotifyURL        string
}

func (c *client) Setup(cfg *ini.File) (err error) {
	err = cfg.Section(ConfSectionWxApplets).MapTo(c)
	if nil != err {
		err = errors.Wrap(err, "解析wxapplets模块的配置失败")
	}

	applets.miniProgram, err = miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     c.AppID,     // 小程序appid
		Secret:    c.AppSecret, // 小程序app secret
		HttpDebug: false,
		Debug:     false,
		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache: nil, // 不传，内部直接使用内存的方式
	})

	applets.payment, err = payment.NewPayment(&payment.UserConfig{
		AppID:       c.AppID,       // 小程序、公众号或者企业微信的appid
		MchID:       c.MchID,       // 商户号 appID
		MchApiV3Key: c.MchApiV3Key, //
		Key:         c.MchApiV2Key,
		CertPath:    c.CertPath,
		KeyPath:     c.KeyPath,
		SerialNo:    c.SerialNo,
		NotifyURL:   c.NotifyURL,
		HttpDebug:   false,
		Debug:       false,
		Log: payment.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Http: payment.Http{
			Timeout: 10,
			BaseURI: "https://api.mch.weixin.qq.com",
		},
		Cache: nil, // 不传，内部直接使用内存的方式
	})

	return
}
