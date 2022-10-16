package redisx

import (
	"fmt"
	"gin_template/modules/component"

	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

var cache client

func init() {
	component.RegComponent(&cache)
}

const (
	ConfSectionRedis = "redis"
)

func (c *client) Setup(config *ini.File) (err error) {
	err = config.Section(ConfSectionRedis).MapTo(c)
	if nil != err {
		err = errors.Wrap(err, "解析redis模块的配置失败")
		return
	}

	c.cli = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       int(c.Db),
	})

	_, err = c.cli.Ping().Result()
	if err != nil {
		err = errors.Wrap(err, "[redisx] ping redis失败")
		return
	}

	return
}
