package mongox

import (
	"context"
	"fmt"
	"gin_template/modules/component"
	"gin_template/modules/logx"
	"time"

	"github.com/go-ini/ini"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

var mgo client

func init() {
	component.RegComponent(&mgo)
}

const (
	ConfSectionMongo = "mongo"
)

func (c *client) Setup(config *ini.File) (err error) {
	err = config.Section(ConfSectionMongo).MapTo(c)
	if nil != err {
		err = errors.Wrap(err, "解析mongo模块的配置失败")
		return
	}

	var url = c.generateUrl()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s", url)).
		SetMaxPoolSize(10).
		SetMonitor(&event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				logx.Infox("[mongox] Started", zap.Reflect("command", evt.Command))
			},
			Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
				logx.Infox("[mongox] Succeeded", zap.Reflect("command", evt.Reply))
			},
			Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
				logx.Errorx("[mongox] Succeeded", zap.Reflect("command", evt.Failure))
			},
		})

	c.cli, err = mongo.Connect(ctx, opts)
	if err != nil {
		err = errors.Wrap(err, "[mongox] 连接mongo失败")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = c.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		err = errors.Wrap(err, "[mongox] ping 失败")
		return
	}

	return
}

func (c *client) generateUrl() string {
	url := c.Host
	if c.User != "" && c.Password != "" {
		url = fmt.Sprintf("%s:%s@%s", c.User, c.Password, url)
	}
	if c.Port != 0 {
		url = fmt.Sprintf("%s:%d", url, c.Port)
	}
	return url
}
