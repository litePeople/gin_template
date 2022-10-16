package alioss

import (
	"gin_template/modules/component"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

var alioss client

func init() {
	component.RegComponent(&alioss)
}

const (
	ConfSectionAliOss = "alioss"
)

type client struct {
	cli             *oss.Client `ini:"-"`
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

func (a *client) Setup(config *ini.File) (err error) {
	err = config.Section(ConfSectionAliOss).MapTo(a)
	if nil != err {
		err = errors.Wrap(err, "解析ali oss模块的配置失败")
	}

	a.cli, err = oss.New(
		a.Endpoint,
		a.AccessKeyID,
		a.AccessKeySecret,
	)
	if err != nil {
		err = errors.Wrap(err, "初始化oss错误")
		return
	}
	return
}
