package centrifugox

import (
	"crypto/tls"
	"gin_template/modules/component"
	"gin_template/modules/utilx"
	"net/http"
	"strings"

	"github.com/centrifugal/gocent"
	"github.com/go-ini/ini"
	"github.com/pkg/errors"
)

var ctfg client

func init() {
	component.RegComponent(&ctfg)
}

const (
	ConfSectionCentrifugo = "centrifugo"
)

type client struct {
	cli                   *gocent.Client `ini:"-"`
	tokenHmacSecretKeyBys []byte         `ini:"-"`
	Addr                  string
	ApiKey                string
	TokenHmacSecretKey    string
}

func (c *client) Setup(config *ini.File) (err error) {
	err = config.Section(ConfSectionCentrifugo).MapTo(c)
	if nil != err {
		err = errors.Wrap(err, "解析centrifugo模块的配置失败")
		return
	}
	insecureSkipVerify := strings.HasPrefix(c.Addr, "https")
	c.tokenHmacSecretKeyBys = utilx.Str2bytes(c.TokenHmacSecretKey)
	c.cli = gocent.New(gocent.Config{
		Addr: c.Addr,
		Key:  c.ApiKey,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
				},
			}},
	})
	return
}
