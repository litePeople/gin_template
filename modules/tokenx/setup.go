package tokenx

import (
	"gin_template/modules/component"

	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

var cli client

func init() {
	component.RegComponent(&cli)
}

const (
	ConfSectionToken = "token"
)

type client struct {
	TokenSigned string
}

func (c *client) Setup(config *ini.File) (err error) {
	err = config.Section(ConfSectionToken).MapTo(c)
	if nil != err {
		err = errors.Wrap(err, "解析token模块的配置失败")
		return
	}
	return
}
